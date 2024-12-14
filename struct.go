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
	"reflect"
	"strconv"
)

// Struct represents any struct.
type Struct struct {
	m map[int]reflect.Value
}

func (structs *Struct) Map() map[interface{}]interface{} {
	return structs.maps(reflect.ValueOf(structs.m))
}

func (structs *Struct) maps(old reflect.Value) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})

	r := old.MapRange()

	for r.Next() {
		switch v := r.Value().Interface().(type) {
		case Struct:
			m[r.Key().Interface()] = v.Map()
		case map[interface{}]interface{}:
			m[r.Key().Interface()] = structs.maps(r.Value())
		case reflect.Value:
			v = Abs[reflect.Value](v)

			if v.Kind() == reflect.Struct && v.Type() == reflect.TypeFor[Struct]() {
				s := v.Interface().(Struct)
				m[r.Key().Interface()] = s.Map()
				continue
			}

			if v.Kind() == reflect.Map {
				m[r.Key().Interface()] = structs.maps(v)
				continue
			}

			m[r.Key().Interface()] = v.Interface()
		default:
			m[r.Key().Interface()] = v
		}
	}

	return m
}

func (structs *Struct) Get(i int) (interface{}, bool) {
	v, ok := structs.m[i]
	if !ok {
		return nil, false
	}

	return v.Interface(), true
}

func (structs *Struct) As(v interface{}) {
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

	structs.ranges(structs.fields(value))
}

func (structs *Struct) fields(value reflect.Value) map[int]reflect.Value {
	fields := make(map[int]reflect.Value)
	typ := value.Type()

	for i := 0; i < value.NumField(); i++ {
		fieldType := typ.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		lookup, ok := fieldType.Tag.Lookup("bin")
		if !ok {
			fields[i+1] = value.Field(i)
			continue
		}

		tag, err := strconv.Atoi(lookup)
		if err != nil {
			continue
		}

		fields[tag] = value.Field(i)
	}

	return fields
}

func (structs *Struct) ranges(fields map[int]reflect.Value) {
	for i, field := range fields {
		m, ok := structs.m[i]
		if !ok {
			continue
		}

		field = Abs[reflect.Value](field)

		case reflect.Map:
			if fields[k].Type() == v.Type() {
				fields[k].Set(v)
				continue
			}

			field := fields[k]
			typ := field.Type()

			m := reflect.MakeMap(reflect.MapOf(typ.Key(), typ.Elem()))
			r := v.MapRange()

			for r.Next() {
				key := r.Key()
				if key.Type() != typ.Key() {
					if key.Kind() == reflect.Interface && key.Elem().Type() == reflect.TypeFor[*Struct]() {
						s := key.Interface().(*Struct)

						key = reflect.New(typ.Elem()).Elem()
						s.ranges(s.fields(key))
					}
				}

				value := r.Value()
				if value.Type() != typ.Elem() {
					if value.Kind() == reflect.Interface && value.Elem().Type() == reflect.TypeFor[*Struct]() {
						s := value.Interface().(*Struct)

						value = reflect.New(typ.Elem()).Elem()
						s.ranges(s.fields(value))
					}
				}

				m.SetMapIndex(key, value)
			}

			fields[k].Set(m)
		switch field.Kind() {
			continue
		case reflect.Struct:
			if fields[k].Kind() == reflect.Interface {
				fields[k].Set(structs.m[k])
				continue
			}

			s, ok := structs.Get(k)
			if !ok {
				continue
			}

			Zero(fields[k])
			s.(*Struct).As(Abs[reflect.Value](fields[k]))
		case reflect.Array, reflect.Slice:
			typ := Abs[reflect.Type](fields[k].Type())
			kind := typ.Kind()

			if kind == reflect.Pointer {
				kind = Abs[reflect.Type](typ).Kind()
			}

			field := fields[k]
			tmp := reflect.New(field.Type()).Elem()

			for i := 0; i < v.Len(); i++ {
				switch kind {
				case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.Interface, reflect.String:
					tmp = reflect.Append(tmp, _struct.ptr(v.Index(i), typ))
				case reflect.Struct:
					ptr := reflect.New(typ).Elem()

					if s, ok := v.Index(i).Interface().(*Struct); ok {
						s.As(ptr)
					}

					tmp = reflect.Append(tmp, ptr)
				default:
				}
			}

			fields[k].Set(tmp)
		default:
			fields[k].Set(v.Convert(fields[k].Type()))
		}
	}
}

func (structs *Struct) ptr(value reflect.Value, typ reflect.Type) reflect.Value {
	value = Abs[reflect.Value](value)

	if value.CanConvert(typ) {
		return value.Convert(typ)
	}

	t := value.Type()

	for tmp := typ; tmp.Kind() == reflect.Ptr; {
		t = reflect.PointerTo(t)
		tmp = tmp.Elem()
	}

	ptr := reflect.New(t).Elem()
	Abs[reflect.Value](ptr).Set(value)

	return ptr
}

func (structs *Struct) Sub(i int, v interface{}) {
	s, ok := structs.Get(i)
	if !ok {
		return
	}

	s.(*Struct).As(&v)
}
