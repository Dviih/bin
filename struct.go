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
	"reflect"
	"strconv"
)

// Struct represents any struct.
type Struct struct {
	m map[int]reflect.Value
}

func (_struct *Struct) Map() map[interface{}]interface{} {
	return _struct._map(reflect.ValueOf(_struct.m))
}

func (_struct *Struct) _map(old reflect.Value) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})

	r := old.MapRange()

	for r.Next() {
		switch v := r.Value().Interface().(type) {
		case Struct:
			m[r.Key().Interface()] = v.Map()
		case map[interface{}]interface{}:
			m[r.Key().Interface()] = _struct._map(r.Value())
		case reflect.Value:
			v = Abs[reflect.Value](v)

			if v.Kind() == reflect.Struct && v.Type() == reflect.TypeFor[Struct]() {
				s := v.Interface().(Struct)
				m[r.Key().Interface()] = s.Map()
				continue
			}

			if v.Kind() == reflect.Map {
				m[r.Key().Interface()] = _struct._map(v)
				continue
			}

			m[r.Key().Interface()] = v.Interface()
		default:
			m[r.Key().Interface()] = v
		}
	}

	return m
}

func (_struct *Struct) Get(i int) (interface{}, bool) {
	v, ok := _struct.m[i]
	if !ok {
		return nil, false
	}

	return v.Interface(), true
}

func (_struct *Struct) As(v interface{}) {
	var value reflect.Value

	if rv, ok := v.(reflect.Value); ok {
		value = rv
	} else {
		value = Abs[reflect.Value](reflect.ValueOf(v))

		if value.Kind() != reflect.Struct || !value.CanSet() {
			return
		}
	}

	if value.Kind() == reflect.Pointer {
		value = Abs[reflect.Value](value)
	}

	_struct.rangeStruct(_struct.fields(value))
}

