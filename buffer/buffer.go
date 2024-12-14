/*
 *     A tiny binary format
 *     Copyright (C) 2024  Dviih
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Affero General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Affero General Public License for more details.
 *
 *     You should have received a copy of the GNU Affero General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 */

package buffer

import (
	"errors"
	"io"
)

const (
	MaxSize = (1 << 31) - 1
)

type Buffer struct {
	data []byte
	read int64

	Max int
}

var (
	InvalidOffset = errors.New("invalid offset")
	InvalidWhence = errors.New("invalid whence")
)

func (buffer *Buffer) Write(data []byte) (int, error) {
	if len(buffer.data)+len(data) > buffer.Max {
		data = data[:buffer.Max-len(buffer.data)]
	}

	buffer.data = append(buffer.data, data...)

	return len(data), nil
}

func (buffer *Buffer) Read(data []byte) (int, error) {
	n := copy(data, buffer.data[buffer.read:buffer.read+int64(len(data))])

	buffer.read += int64(n)

	if len(data) > n {
		return n, io.EOF
	}

	return n, nil
}

func (buffer *Buffer) ReadByte() (byte, error) {
	if buffer.read >= int64(len(buffer.data)) {
		return 0, io.EOF
	}

	buffer.read++
	return buffer.data[buffer.read-1], nil
}

func (buffer *Buffer) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		if buffer.read < offset {
			return 0, InvalidOffset
		}

		buffer.read = offset
	case io.SeekCurrent:
		if buffer.read+offset > int64(len(buffer.data)) {
			return 0, InvalidOffset
		}

		buffer.read += offset
	case io.SeekEnd:
		if buffer.read+offset > int64(len(buffer.data)) {
			return 0, InvalidOffset
		}

		buffer.read = int64(len(buffer.data)) - offset
	default:
		return 0, InvalidWhence
	}

	return buffer.read, nil
}

func (buffer *Buffer) Data() []byte {
	return buffer.data[:]
}

func (buffer *Buffer) Len() int {
	return len(buffer.data)
}

