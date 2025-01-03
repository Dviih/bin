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
	"errors"
	"github.com/Dviih/bin/buffer"
	"reflect"
)

var (
	Invalid              = errors.New("invalid value")
	CantSet              = errors.New("can't set")
	TypeMustBeComparable = errors.New("type must be comparable")
	unexpectedBehaviour  = errors.New("this is a very unexpected behaviour")
)

func Value(v interface{}) reflect.Value {
	if rv, ok := v.(reflect.Value); ok {
		return rv
	} else {
		return Abs[reflect.Value](reflect.ValueOf(v))
	}
}

func Zero(value reflect.Value) {
	for value.Kind() == reflect.Pointer {
		if !value.CanSet() {
			break
		}

		value.Set(reflect.New(value.Type().Elem()))
		value = value.Elem()
	}
}

func Abs[T interface{}](t interface{}) T {
	switch t := t.(type) {
	case reflect.Value:
		for {
			switch t.Kind() {
			case reflect.Pointer, reflect.Interface:
				if t.IsZero() {
					Zero(t)
				}

				elem := t.Elem()

				if elem.Kind() == reflect.Invalid {
					return (interface{})(t).(T)
				}

				t = elem
			default:
				return (interface{})(t).(T)
			}
		}
	case reflect.Type:
		for {
			switch t.Kind() {
			case reflect.Pointer:
				t = t.Elem()
			default:
				return t.(T)
			}
		}
	}

	return t.(T)
}

func KeyElem(value reflect.Value) (reflect.Type, reflect.Type) {
	t := Abs[reflect.Value](value).Type()

	for {
		switch t.Kind() {
		case reflect.Array, reflect.Slice:
			t = t.Elem()
		case reflect.Map:
			k, v := t.Key(), t.Elem()

			_, k = KeyElem(reflect.New(k))
			_, v = KeyElem(reflect.New(v))

			return k, v
		default:
			return nil, t
		}
	}
}

func Marshal(v interface{}) ([]byte, error) {
	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(v); err != nil {
		return nil, err
	}

	return b.Data(), nil
}

func Unmarshal[T interface{}](data []byte) (T, error) {
	var t T

	if err := NewDecoder(buffer.From(data)).Decode(&t); err != nil {
		var zero T
		return zero, err
	}

	return t, nil
}

func UnmarshalAs[T interface{}](data []byte) (T, error) {
	i, err := Unmarshal[interface{}](data)
	if err != nil {
		var zero T
		return zero, err
	}

	return As[T](i), nil
}
