/*
 *     A tiny binary format
 *     Copyright (C) 2025  Dviih
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
	"io"
	"reflect"
	"unsafe"
)

type Integer = interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func VarIntIn[T Integer](writer io.Writer, t T) error {
	b := make([]byte, 0, 10)

	for int(t) >= 0x80 {
		b = append(b, byte(t)|0x80)
		t >>= 7
	}

	b = append(b, byte(t))

	if _, err := writer.Write(b); err != nil {
		return err
	}

	return nil
}

func VarIntOut[T Integer](reader io.Reader) (T, error) {
	var br func() (byte, error)

	if rbr, ok := reader.(io.ByteReader); ok {
		br = rbr.ReadByte
	} else {
		br = func() (byte, error) {
			b := [1]byte{}

			n, err := reader.Read(b[:])
			if err != nil {
				return 0, err
			}

			if n != 1 {
				return 0, io.EOF
			}

			return b[0], nil
		}
	}

	var t T
	var p uint64

	for i := 0; i < 10; i++ {
		b, err := br()
		if err != nil {
			return 0, err
		}

		if b < 0x80 {
			if i == 9 && b > 1 {
				return 0, io.EOF
			}

			return t | T(b)<<p, nil
		}

		t |= T(b&0x7f) << p
		p += 7
	}

	return 0, io.EOF
}

	}

	return T(t), nil
}
