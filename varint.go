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

package bin

import (
	"encoding/binary"
	"io"
)

type Integer = interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func VarIntIn[T Integer](writer io.Writer, t T) error {
	if _, err := writer.Write(binary.AppendUvarint(nil, uint64(t))); err != nil {
		return err
	}

	return nil
}

func VarIntOut[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64](reader io.ByteReader) (T, error) {
	t, err := binary.ReadUvarint(reader)
	if err != nil {
		return 0, err
	}

	return T(t), nil
}
