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
	"bytes"
	"errors"
	"reflect"
)

var (
	Invalid             = errors.New("invalid value")
	CantSet             = errors.New("can't set")
	unexpectedBehaviour = errors.New("this is a very unexpected behaviour")
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

func Marshal(v interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := NewEncoder(buffer)

	if err := encoder.Encode(v); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
