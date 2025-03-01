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

//go:build !dviih_bin_kind_encoding

package bin

import (
	"encoding"
	"github.com/Dviih/bin/kind"
	"reflect"
)

// This init function is responsible to add handlers
// for `encoding.BinaryMarshaler`, `encoding.BinaryUnmarshaler`,
// `encoding.TextMarshaler` and `encoding.TextUnmarshaler`
func init() {
	b := kind.NewHandler(
		func(encoder kind.Encoder, value reflect.Value) error {
			out := kind.Call(value, "MarshalBinary")
			if !out[1].IsNil() {
				return out[1].Interface().(error)
			}

			return encoder.Encode(out[0].Interface().([]byte))
		},
		func(decoder kind.Decoder, value reflect.Value) error {
			var data []byte

			if err := decoder.Decode(&data); err != nil {
				return err
			}

			out := kind.Call(value, "UnmarshalBinary", reflect.ValueOf(data))
			if !out[0].IsNil() {
				return out[0].Interface().(error)
			}

			return nil
		},
	)

	t := kind.NewHandler(
		func(encoder kind.Encoder, value reflect.Value) error {
			tm := value.MethodByName("MarshalText")
			out := tm.Call(nil)

			if !out[1].IsNil() {
				return out[1].Interface().(error)
			}

			return encoder.Encode(out[0].Interface().([]byte))
		},
		func(decoder kind.Decoder, value reflect.Value) error {
			var data []byte

			if err := decoder.Decode(&data); err != nil {
				return err
			}

			ut := value.MethodByName("UnmarshalText")

			out := ut.Call([]reflect.Value{reflect.ValueOf(data)})
			if !out[0].IsNil() {
				return out[0].Interface().(error)
			}

			return nil
		},
	)

	register(65, reflect.TypeFor[encoding.BinaryMarshaler](), b)
	mkind.Alias(65, reflect.TypeFor[encoding.BinaryUnmarshaler]())
}
