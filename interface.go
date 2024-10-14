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
)

func Interface(v interface{}) reflect.Value {
	if v == nil {
		return reflect.New(reflect.TypeFor[interface{}]()).Elem()
	}

	ptr := reflect.New(reflect.TypeFor[interface{}]()).Elem()
	ptr.Set(_interface(Value(v)))

	return ptr
}

func _interface(value reflect.Value) reflect.Value {
	switch value.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice:
		return value
	case reflect.Struct:
		var fields []reflect.StructField
		var values []reflect.Value

		typ := value.Type()

		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			fieldType := typ.Field(i)

			if !fieldType.IsExported() {
				continue
			}

			if lookup, ok := fieldType.Tag.Lookup("bin"); ok && lookup == "-" {
				continue
			}

			fields = append(fields, fieldType)
			values = append(values, field)
		}

		tmp := reflect.New(reflect.StructOf(fields)).Elem()

		for i, v := range values {
			tmp.Field(i).Set(v)
		}

		return tmp
	default:
		return value.Convert(reflect.TypeFor[interface{}]())
	}
}
