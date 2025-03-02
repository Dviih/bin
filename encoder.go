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
	"strconv"
)

type Encoder struct {
	writer io.Writer
}

func (encoder *Encoder) Encode(v interface{}) error {
	value := Value(v)

	if v != nil {
		found, err := mkind.Run(value.Type(), encoder, value)
		if err != nil {
			return err
		}

		if found {
			return nil
		}
	}

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
		return encoder.Encode(floatToBits(float32(value.Float())))
	case reflect.Float64:
		return encoder.Encode(floatToBits(value.Float()))
	case reflect.Complex64:
		c := complex64(value.Complex())

		if err := encoder.Encode(floatToBits(real(c))); err != nil {
			return err
		}

		return encoder.Encode(floatToBits(imag(c)))
	case reflect.Complex128:
		c := value.Complex()

		if err := encoder.Encode(floatToBits(real(c))); err != nil {
			return err
		}

		return encoder.Encode(floatToBits(imag(c)))
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

		if n, ok := mkind.Has(value.Type()); ok {
			if err := encoder.Encode(n); err != nil {
				return err
			}

			_, err := mkind.Run(n, encoder, value)
			return err
		}

		switch value.Kind() {
		case reflect.Array, reflect.Slice:
			_, elem := KeyElem(value)

			switch Abs[reflect.Type](elem).Kind() {
			case reflect.Struct:
				if err := encoder.getType(reflect.New(reflect.TypeFor[[]interface{}]()).Elem()); err != nil {
					return err
				}

				if err := encoder.Encode(value.Len()); err != nil {
					return err
				}

				for i := 0; i < value.Len(); i++ {
					if err := encoder.Encode(interfaces(value.Index(i))); err != nil {
						return err
					}
				}

				return nil
			default:
				if err := encoder.getType(value); err != nil {
					return err
				}
			}
		case reflect.Map:
			key, elem := KeyElem(value)

			if !key.Comparable() {
				return TypeMustBeComparable
			}

			switch Abs[reflect.Type](elem).Kind() {
			case reflect.Struct:
				if err := encoder.getType(reflect.New(reflect.MapOf(key, reflect.TypeFor[interface{}]())).Elem()); err != nil {
					return err
				}

				if err := encoder.Encode(value.Len()); err != nil {
					return err
				}

				m := value.MapRange()

				for m.Next() {
					if err := encoder.Encode(m.Key()); err != nil {
						return err
					}

					if err := encoder.Encode(interfaces(m.Value())); err != nil {
						return err
					}
				}

				return nil
			default:
				if err := encoder.getType(value); err != nil {
					return err
				}
			}

		case reflect.Struct:
			if err := encoder.getType(value); err != nil {
				return err
			}

			return encoder.structs(value, true)
		default:
			if err := encoder.getType(value); err != nil {
				return err
			}
		}

		return encoder.Encode(value)
	case reflect.Map:
		if !value.Type().Key().Comparable() {
			return TypeMustBeComparable
		}

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

		if _, err := encoder.writer.Write([]byte(value.String())); err != nil {
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
		field := value.Field(i)
		if kind && field.IsZero() {
			continue
		}

		ft := t.Field(i)

		if !ft.IsExported() {
			continue
		}

		kind := kind
		if !kind && ft.Type.Kind() == reflect.Interface {
			kind = true
		}

		tag := i + 1

		if lookup, ok := ft.Tag.Lookup("bin"); ok {
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

		if field.IsZero() {
			if err := encoder.Encode(0); err != nil {
				return err
			}

			continue
		}

		lf, _ := mkind.Load(field.Type())
		if lf != 0 {
			if kind {
				if err := encoder.Encode(lf); err != nil {
					return err
				}
			}

			if _, err := mkind.Run(lf, encoder, field); err != nil {
				return err
			}

			continue
		} else if kind {
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

		if err := encoder.Encode(Abs[reflect.Type](dt).Kind()); err != nil {
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

		kind := Abs[reflect.Type](dt).Kind()

		if err := encoder.Encode(kind); err != nil {
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
		n := value.Type().NumField()

		for i := 0; i < value.Type().NumField(); i++ {
			if value.Field(i).IsZero() {
				n--
			}
		}

		if err := encoder.Encode(n); err != nil {
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
