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
)

type Decoder struct {
	reader io.Reader
}

func (decoder *Decoder) Decode(v interface{}) error {
	value := Value(v)

	if !value.CanSet() {
		return CantSet
	}

	if value.Type() == reflect.TypeFor[*Struct]() {
		Zero(value)
		return decoder.structs(value)
	}

	found, err := mkind.Run(value.Type(), decoder, value)
	if err != nil {
		return err
	}

	if found {
		return nil
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
		b := [1]byte{}

		n, err := decoder.reader.Read(b[:])
		if err != nil {
			return err
		}

		if n != 1 {
			return io.EOF
		}

		if b[0] == 255 {
			value.Set(reflect.ValueOf(true))
			return nil
		}

		value.Set(reflect.ValueOf(false))
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := VarIntOut[int64](decoder.reader)
		if err != nil {
			return err
		}

		value.SetInt(n)
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := VarIntOut[uint64](decoder.reader)
		if err != nil {
			return err
		}

		value.SetUint(n)
		return nil
	case reflect.Float32:
		n, err := VarIntOut[uint32](decoder.reader)
		if err != nil {
			return err
		}

		value.SetFloat(floatFromBits(n))
		return nil
	case reflect.Float64:
		n, err := VarIntOut[uint64](decoder.reader)
		if err != nil {
			return err
		}

		value.SetFloat(floatFromBits(n))
		return nil
	case reflect.Complex64:
		r, err := VarIntOut[uint32](decoder.reader)
		if err != nil {
			return err
		}

		i, err := VarIntOut[uint32](decoder.reader)
		if err != nil {
			return err
		}

		value.SetComplex(complex(floatFromBits(r), floatFromBits(i)))
		return nil
	case reflect.Complex128:
		r, err := VarIntOut[uint64](decoder.reader)
		if err != nil {
			return err
		}

		i, err := VarIntOut[uint64](decoder.reader)
		if err != nil {
			return err
		}

		value.SetComplex(complex(floatFromBits(r), floatFromBits(i)))
		return nil
	case reflect.Array:
		for i := 0; i < value.Len(); i++ {
			if err := decoder.Decode(value.Index(i)); err != nil && err != io.EOF {
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

		found, t, err := decoder.getType()
		if err != nil {
			return err
		}

		if t == nil {
			b := [1]byte{}

			if _, err := decoder.reader.Read(b[:]); err != nil && err != io.EOF {
				return err
			}

			if b[0] != 0 {
				return unexpectedBehavior
			}

			return nil
		}

		if found {
			ptr := reflect.New(t)
			if _, err = mkind.Run(t, decoder, ptr.Elem()); err != nil {
				return err
			}

			value.Set(ptr.Elem())
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
		size, err := VarIntOut[int](decoder.reader)
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
		size, err := VarIntOut[int](decoder.reader)
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
		size, err := VarIntOut[int](decoder.reader)
		if err != nil {
			return err
		}

		data := make([]byte, size)

		if _, err = decoder.reader.Read(data); err != nil {
			return err
		}

		value.SetString(string(data))
		return nil
	case reflect.Struct:
		fields := (&Struct{}).fields(value)

		for i := 0; i < len(fields); i++ {
			tag, err := VarIntOut[int](decoder.reader)
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
	size, err := VarIntOut[int](decoder.reader)
	if err != nil {
		return err
	}

	s := &Struct{
		m: make(map[int]reflect.Value),
	}

	value.Set(reflect.ValueOf(s))

	for i := 0; i < size; i++ {
		tag, err := VarIntOut[int](decoder.reader)
		if err != nil {
			return err
		}

		found, t, err := decoder.getType()
		if err != nil {
			return err
		}

		var ptr reflect.Value

		if found {
			ptr = reflect.New(t).Elem()

			if _, err = mkind.Run(t, decoder, ptr); err != nil {
				return err
			}

			s.m[tag] = ptr
			continue
		}

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

func (decoder *Decoder) getType() (bool, reflect.Type, error) {
	kind, err := VarIntOut[int](decoder.reader)
	if err != nil {
		return false, nil, err
	}

	switch reflect.Kind(kind) {
	case reflect.Invalid:
		return false, nil, nil
	case reflect.Bool:
		return false, reflect.TypeFor[bool](), nil
	case reflect.Int:
		return false, reflect.TypeFor[int](), nil
	case reflect.Int8:
		return false, reflect.TypeFor[int8](), nil
	case reflect.Int16:
		return false, reflect.TypeFor[int16](), nil
	case reflect.Int32:
		return false, reflect.TypeFor[int32](), nil
	case reflect.Int64:
		return false, reflect.TypeFor[int64](), nil
	case reflect.Uint:
		return false, reflect.TypeFor[uint](), nil
	case reflect.Uint8:
		return false, reflect.TypeFor[uint8](), nil
	case reflect.Uint16:
		return false, reflect.TypeFor[uint16](), nil
	case reflect.Uint32:
		return false, reflect.TypeFor[uint32](), nil
	case reflect.Uint64:
		return false, reflect.TypeFor[uint64](), nil
	case reflect.Uintptr:
		return false, reflect.TypeFor[uintptr](), nil
	case reflect.Float32:
		return false, reflect.TypeFor[float32](), nil
	case reflect.Float64:
		return false, reflect.TypeFor[float64](), nil
	case reflect.Complex64:
		return false, reflect.TypeFor[complex64](), nil
	case reflect.Complex128:
		return false, reflect.TypeFor[complex128](), nil
	case reflect.Interface:
		return false, reflect.TypeFor[interface{}](), nil
	case reflect.String:
		return false, reflect.TypeFor[string](), nil
	case reflect.Array:
		d, err := VarIntOut[int](decoder.reader)
		if err != nil {
			return false, nil, err
		}

		var mixed bool
		if err = decoder.Decode(&mixed); err != nil {
			return false, nil, err
		}

		var di []int
		for i := 0; i < d; i++ {
			n, err := VarIntOut[int](decoder.reader)
			if err != nil {
				return false, nil, err
			}

			di = append(di, n)
		}

		found, t, err := decoder.getType()
		if err != nil {
			return found, nil, err
		}

		return found, fromDepth(t, d, di), nil
	case reflect.Slice:
		d, err := VarIntOut[int](decoder.reader)
		if err != nil {
			return false, nil, err
		}

		var mixed bool
		if err = decoder.Decode(&mixed); err != nil {
			return false, nil, err
		}

		var di []int

		if mixed {
			for i := 0; i < d; i++ {
				n, err := VarIntOut[int](decoder.reader)
				if err != nil {
					return false, nil, err
				}

				di = append(di, n)
			}
		}

		found, t, err := decoder.getType()
		if err != nil {
			return found, nil, err
		}

		return found, fromDepth(t, d, di), nil
	case reflect.Map:
		found, key, err := decoder.getType()
		if err != nil {
			return found, nil, err
		}

		found2, value, err := decoder.getType()
		if err != nil {
			return found2, nil, err
		}

		if !found && found2 {
			found = true
		}

		return found, reflect.MapOf(key, value), nil
	case reflect.Struct:
		return false, reflect.TypeFor[*Struct](), nil
	case reflect.Chan, reflect.Func, reflect.Pointer, reflect.UnsafePointer:
		return false, nil, nil
	default:
		_, lt := mkind.Load(kind)
		if lt != nil {
			return true, lt, nil
		}
	}

	return false, nil, nil
}
