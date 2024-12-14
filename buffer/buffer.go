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

