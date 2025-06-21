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
	"unsafe"
)

func realType(p reflect.Type) reflect.Type {
	switch p.Kind() {
	case reflect.Invalid:
		return reflect.TypeOf(nil)
	case reflect.Bool:
		return reflect.TypeFor[bool]()
	case reflect.Int:
		return reflect.TypeFor[int]()
	case reflect.Int8:
		return reflect.TypeFor[int8]()
	case reflect.Int16:
		return reflect.TypeFor[int16]()
	case reflect.Int32:
		return reflect.TypeFor[int32]()
	case reflect.Int64:
		return reflect.TypeFor[int64]()
	case reflect.Uint:
		return reflect.TypeFor[int]()
	case reflect.Uint8:
		return reflect.TypeFor[int8]()
	case reflect.Uint16:
		return reflect.TypeFor[uint16]()
	case reflect.Uint32:
		return reflect.TypeFor[uint32]()
	case reflect.Uint64:
		return reflect.TypeFor[uint64]()
	case reflect.Uintptr:
		return reflect.TypeFor[uintptr]()
	case reflect.Float32:
		return reflect.TypeFor[float32]()
	case reflect.Float64:
		return reflect.TypeFor[float64]()
	case reflect.Complex64:
		return reflect.TypeFor[complex64]()
	case reflect.Complex128:
		return reflect.TypeFor[complex128]()
	case reflect.Array:
		return reflect.ArrayOf(p.Len(), realType(p.Elem()))
	case reflect.Chan:
		return reflect.ChanOf(p.ChanDir(), realType(p.Elem()))
	case reflect.Func:
		var in, out []reflect.Type

		for i := 0; i < p.NumIn(); i++ {
			in = append(in, realType(p.In(i)))
		}

		for i := 0; i < p.NumOut(); i++ {
			out = append(out, realType(p.Out(i)))
		}

		return reflect.FuncOf(in, out, p.IsVariadic())
	case reflect.Interface:
		return nil
	case reflect.Map:
		return reflect.MapOf(realType(p.Key()), realType(p.Elem()))
	case reflect.Pointer:
		return reflect.TypeFor[uintptr]()
	case reflect.Slice:
		return reflect.SliceOf(realType(p.Elem()))
	case reflect.String:
		return reflect.TypeFor[string]()
	case reflect.Struct:
		var fields []reflect.StructField

		for i := 0; i < p.NumField(); i++ {
			fields = append(fields, p.Field(i))
		}

		return reflect.StructOf(fields)
	case reflect.UnsafePointer:
		return reflect.TypeFor[unsafe.Pointer]()
	default:
		return nil
	}
}

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

var EncodingText = NewHandler(
	func(encoder Encoder, value reflect.Value) error {
		out := Call(value, "MarshalText")
		if !out[1].IsNil() {
			return out[1].Interface().(error)
		}

		if err := encoder.Encode(out[0].Len()); err != nil {
			return err
		}

		return encoder.Encode(out[0].Interface().([]byte))
	},
	func(decoder Decoder, value reflect.Value) error {
		var data []byte

		if err := decoder.Decode(&data); err != nil {
			return err
		}

		out := Call(value, "UnmarshalText", reflect.ValueOf(data))
		if !out[0].IsNil() {
			return out[0].Interface().(error)
		}

		return nil
	},
)

// This file includes the handler for `encoding/gob`,
// since bin acts as an alternative to gob, it is only
// used for some packages that implement it but not
// the standard.

var Gob = NewHandler(
	func(encoder Encoder, value reflect.Value) error {
		i := Call(value, "GobEncode")
		if !i[1].IsNil() {
			return i[1].Interface().(error)
		}

		return encoder.Encode(i[0].Interface().([]byte))
	},
	func(decoder Decoder, value reflect.Value) error {
		var data []byte

		if err := decoder.Decode(&data); err != nil {
			return err
		}

		if i := Call(value, "GobDecode", reflect.ValueOf(data)); !i[0].IsNil() {
			return i[0].Interface().(error)
		}

		return nil
	},
)
