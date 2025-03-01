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

package kind

import (
	"reflect"
)

var EncodingBinary = NewHandler(
	func(encoder Encoder, value reflect.Value) error {
		out := Call(value, "MarshalBinary")
		if !out[1].IsNil() {
			return out[1].Interface().(error)
		}

		return encoder.Encode(out[0].Interface().([]byte))
	},
	func(decoder Decoder, value reflect.Value) error {
		var data []byte

		if err := decoder.Decode(&data); err != nil {
			return err
		}

		if len(data) == 0 {
			return nil
		}

		out := Call(value, "UnmarshalBinary", reflect.ValueOf(data))
		if !out[0].IsNil() {
			return out[0].Interface().(error)
		}

		return nil
	},
)

