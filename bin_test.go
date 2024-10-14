/*
 *     A tiny format for using binary data
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

package bin

import (
	"reflect"
	"testing"
)

type stream struct {
	Data []byte
	read int
}

func (stream *stream) Write(data []byte) (int, error) {
	stream.Data = append(stream.Data, data...)
	return len(data), nil
}

func (stream *stream) Read(data []byte) (int, error) {
	i := copy(data, stream.Data[stream.read:stream.read+len(data)])
	stream.read += len(data)

	return i, nil
}

