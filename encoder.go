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
	"bytes"
	"io"
	"math"
	"reflect"
	"strconv"
)

type Encoder struct {
	writer io.Writer
}

func (encoder *Encoder) Encode(v interface{}) error {
	value := Value(v)

	switch value.Kind() {
	case reflect.Invalid, reflect.Uintptr, reflect.UnsafePointer:
		if v == nil {
			return encoder.Encode(0)
		}

		return Invalid
	case reflect.Bool:
		if value.Bool() {
			_, err := encoder.writer.Write([]byte{255})
			return err
		}

		return encoder.Encode(0)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if err := VarIntIn(encoder.writer, value.Int()); err != nil {
			return err
		}
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if err := VarIntIn(encoder.writer, value.Uint()); err != nil {
			return err
		}

		return nil
	case reflect.Float32:
		return encoder.Encode(math.Float32bits(float32(value.Float())))
	case reflect.Float64:
		return encoder.Encode(math.Float64bits(value.Float()))
	case reflect.Complex64:
		c := complex64(value.Complex())

		if err := encoder.Encode(math.Float32bits(real(c))); err != nil {
			return err
		}

		return encoder.Encode(math.Float32bits(imag(c)))
	case reflect.Complex128:
		c := value.Complex()

		if err := encoder.Encode(math.Float64bits(real(c))); err != nil {
			return err
		}

		return encoder.Encode(math.Float64bits(imag(c)))
	case reflect.Array:
		for i := 0; i < value.Len(); i++ {
			if err := encoder.Encode(value.Index(i)); err != nil {
				return err
			}
		}
	case reflect.Chan, reflect.Func:
		// Channels and Functions aren't supported.
		return nil
	case reflect.Interface:
		if value.IsNil() {
			if err := encoder.Encode(reflect.Invalid); err != nil {
				return err
			}

			if err := encoder.Encode(nil); err != nil {
				return err
			}

			return nil
		}

		value = Abs[reflect.Value](value)

		if err := encoder.getType(value); err != nil {
			return err
		}

		switch value.Kind() {
		case reflect.Array, reflect.Slice, reflect.Map:
			switch value.Type().Elem().Kind() {
			case reflect.Struct:
				for i := 0; i < value.Len(); i++ {
					if err := encoder.Encode(_interface(value.Index(i))); err != nil {
						return err
					}
				}

				return nil
			default:
			}
		case reflect.Struct:
			return encoder.structs(value, true)
		default:
		}

		return encoder.Encode(value)
	case reflect.Map:
		if err := encoder.Encode(value.Len()); err != nil {
			return err
		}

		m := value.MapRange()

		for m.Next() {
			if err := encoder.Encode(m.Key()); err != nil {
				return err
			}

			if err := encoder.Encode(m.Value()); err != nil {
				return err
			}
		}
	case reflect.Pointer:
		for value.Kind() == reflect.Pointer {
			value = value.Elem()
		}

		return encoder.Encode(value)
	case reflect.Slice:
		if err := encoder.Encode(value.Len()); err != nil {
			return err
		}

		for i := 0; i < value.Len(); i++ {
			if err := encoder.Encode(value.Index(i)); err != nil {
				return err
			}
		}
	case reflect.String:
		if err := encoder.Encode(value.Len()); err != nil {
			return err
		}

		if _, err := io.Copy(encoder.writer, bytes.NewBufferString(value.String())); err != nil {
			return err
		}
	case reflect.Struct:
		return encoder.structs(value, false)
	}

	return nil
}

func (encoder *Encoder) structs(value reflect.Value, kind bool) error {
	t := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := Abs[reflect.Value](value.Field(i))
		if field.IsZero() {
			continue
		}

		fieldType := t.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		tag := i + 1

		if lookup, ok := fieldType.Tag.Lookup("bin"); ok {
			if lookup == "-" {
				continue
			}

			n, err := strconv.Atoi(lookup)
			if err != nil {
				return err
			}

			tag = n
		}

		if err := encoder.Encode(tag); err != nil {
			return err
		}

		field := Abs[reflect.Value](value.Field(i))

		if kind {
			if err := encoder.Encode(Interface(field.Interface())); err != nil {
				return err
			}

			continue
		}

		if err := encoder.Encode(field); err != nil {
			return err
		}
	}

	return nil
}

func (encoder *Encoder) getType(value reflect.Value) error {
	if err := encoder.Encode(value.Type().Kind()); err != nil {
		return err
	}

	switch value.Type().Kind() {
	case reflect.Array:
		dt, d, mixed, di := depth(value)

		if err := encoder.Encode(d); err != nil {
			return err
		}

		if err := encoder.Encode(mixed); err != nil {
			return err
		}

		for i := 0; i < len(di); i++ {
			if err := encoder.Encode(di[i]); err != nil {
				return err
			}
		}

		if err := encoder.Encode(dt.Kind()); err != nil {
			return err
		}

		return nil
	case reflect.Slice:
		dt, d, mixed, di := depth(value)

		if err := encoder.Encode(d); err != nil {
			return err
		}

		if err := encoder.Encode(mixed); err != nil {
			return err
		}

		if mixed {
			for i := 0; i < len(di); i++ {
				if err := encoder.Encode(di[i]); err != nil {
					return err
				}
			}
		}

		if err := encoder.Encode(dt.Kind()); err != nil {
			return err
		}

		return nil
	case reflect.Map:
		if err := encoder.getType(reflect.New(value.Type().Key()).Elem()); err != nil {
			return err
		}

		if err := encoder.getType(reflect.New(value.Type().Elem()).Elem()); err != nil {
			return err
		}

		return nil
	case reflect.Struct:
		if err := encoder.Encode(value.Type().NumField()); err != nil {
			return nil
		}

		return nil
	default:
		return nil
	}
}

func NewEncoder(writer io.Writer) *Encoder {
	return &Encoder{writer: writer}
}
