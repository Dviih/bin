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
	"io"
	"math"
	"reflect"
)

type Decoder struct {
	reader io.Reader
}

func (decoder *Decoder) ReadByte() (byte, error) {
	data := make([]byte, 1)
	_, err := decoder.reader.Read(data)

	return data[0], err
}

func (decoder *Decoder) Decode(v interface{}) error {
	value := Value(v)

	if !value.CanSet() {
		return CantSet
	}

	if v == nil {
		value.SetZero()
		return nil
	}

	switch value.Kind() {
	case reflect.Invalid, reflect.Uintptr, reflect.UnsafePointer:
		value.SetZero()
		return nil
	case reflect.Bool:
		b, err := decoder.ReadByte()
		if err != nil {
			return err
		}

		if b == 255 {
			value.Set(reflect.ValueOf(true))
			return nil
		}

		value.Set(reflect.ValueOf(false))
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := VarIntOut[int64](decoder)
		if err != nil {
			return err
		}

		value.SetInt(n)
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := VarIntOut[uint64](decoder)
		if err != nil {
			return err
		}

		value.SetUint(n)
		return nil
	case reflect.Float32:
		n, err := VarIntOut[uint32](decoder)
		if err != nil {
			return err
		}

		value.Set(reflect.ValueOf(math.Float32frombits(n)))
		return nil
	case reflect.Float64:
		n, err := VarIntOut[uint64](decoder)
		if err != nil {
			return err
		}

		value.Set(reflect.ValueOf(math.Float64frombits(n)))
		return nil
	case reflect.Complex64:
		r, err := VarIntOut[uint32](decoder)
		if err != nil {
			return err
		}

		i, err := VarIntOut[uint32](decoder)
		if err != nil {
			return err
		}

		value.Set(reflect.ValueOf(complex(math.Float32frombits(r), math.Float32frombits(i))))
		return nil
	case reflect.Complex128:
		r, err := VarIntOut[uint64](decoder)
		if err != nil {
			return err
		}

		i, err := VarIntOut[uint64](decoder)
		if err != nil {
			return err
		}

		value.Set(reflect.ValueOf(complex(math.Float64frombits(r), math.Float64frombits(i))))
		return nil
	case reflect.Array:
		for i := 0; i < value.Len(); i++ {
			if err := decoder.Decode(value.Index(i)); err != nil {
				return err
			}
		}

		return nil
	case reflect.Chan, reflect.Func:
		return nil
	case reflect.Interface:
		if value.Kind() == reflect.Invalid {
			return nil
		}

		t, err := decoder.getType()
		if err != nil {
			return err
		}

		if t == nil {
			return nil
		}

		if t.Kind() == reflect.Struct {
			return decoder.structs(value)
		}

		ptr := reflect.New(t).Elem()

		if err = decoder.Decode(ptr); err != nil {
			return err
		}

		value.Set(ptr)
		return nil
	case reflect.Map:
		size, err := VarIntOut[int](decoder)
		if err != nil {
			return err
		}

		value.Set(reflect.MakeMapWithSize(value.Type(), size))

		keyType := value.Type().Key()
		valueType := value.Type().Elem()

		for i := 0; i < size; i++ {
			mk := reflect.New(keyType).Elem()
			if err = decoder.Decode(mk); err != nil {
				return err
			}

			mv := reflect.New(valueType).Elem()
			if err = decoder.Decode(mv); err != nil {
				return err
			}

			value.SetMapIndex(mk, mv)
		}

		return nil
	case reflect.Pointer:
		Zero(value)
		for value.Kind() == reflect.Pointer {
			value = value.Elem()
		}

		return decoder.Decode(value)
	case reflect.Slice:
		size, err := VarIntOut[int](decoder)
		if err != nil {
			return err
		}

		value.Set(reflect.MakeSlice(value.Type(), size, size))

		for i := 0; i < size; i++ {
			if err = decoder.Decode(value.Index(i)); err != nil {
				return err
			}
		}

		return nil
	case reflect.String:
		size, err := VarIntOut[int](decoder)
		if err != nil {
			return err
		}

		data := make([]byte, size)

		if _, err = decoder.reader.Read(data); err != nil {
			return err
		}

		value.Set(reflect.ValueOf(string(data)))
		return nil
	case reflect.Struct:
		fields := (&Struct{}).fields(value)

		for i := 0; i < len(fields); i++ {
			tag, err := VarIntOut[int](decoder)
			if err != nil {
				return err
			}

			field, ok := fields[tag]
			if !ok {
				continue
			}

			Zero(field)
			if err = decoder.Decode(field); err != nil {
				return err
			}
		}

		return nil
	}

	return Invalid
}

func (decoder *Decoder) structs(value reflect.Value) error {
	size, err := VarIntOut[int](decoder)
	if err != nil {
		return err
	}

	s := &Struct{
		m: make(map[int]reflect.Value),
	}

	value.Set(reflect.ValueOf(s))

	for i := 0; i < size; i++ {
		tag, err := VarIntOut[int](decoder)
		if err != nil {
			return err
		}

		t, err := decoder.getType()
		if err != nil {
			return err
		}

		var ptr reflect.Value

		if t.Kind() == reflect.Struct {
			ptr = reflect.New(reflect.TypeFor[interface{}]()).Elem()
			if err = decoder.structs(ptr); err != nil {
				return err
			}
		} else {
			ptr = reflect.New(t).Elem()
			if err = decoder.Decode(ptr); err != nil {
				return err
			}
		}

		s.m[tag] = ptr
	}

	return nil
}

func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{
		reader: reader,
	}
}

	switch reflect.Kind(kind) {
	case reflect.Bool:
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.Uintptr:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Complex64:
	case reflect.Complex128:
	case reflect.String:
	}

}
