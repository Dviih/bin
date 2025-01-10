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
		switch v := Abs[reflect.Value](r.Value()).Interface().(type) {
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
		case []interface{}:
			m[r.Key()] = structs.arrays(Abs[reflect.Value](r.Value()))
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

		switch field.Kind() {
		case reflect.Invalid, reflect.Uintptr, reflect.Pointer, reflect.UnsafePointer, reflect.Chan, reflect.Func:
			continue
		case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String, reflect.Array, reflect.Slice, reflect.Map:
			field.Set(structs.convert(field.Type(), m))
		case reflect.Interface:
			field.Set(m)
		case reflect.Struct:
			s, ok := m.Interface().(*Struct)
			if !ok {
				continue
			}

			s.As(field)
		}
	}
}

func (structs *Struct) ptr(typ reflect.Type, value reflect.Value) reflect.Value {
	t := value.Type()

	for tmp := typ; tmp.Kind() == reflect.Ptr; {
		t = reflect.PointerTo(t)
		tmp = tmp.Elem()
	}

	ptr := reflect.New(t).Elem()
	Abs[reflect.Value](ptr).Set(value)

	return ptr
}

func (structs *Struct) convert(t reflect.Type, value reflect.Value) reflect.Value {
	if value.CanConvert(t) {
		return value.Convert(t)
	}

	if Abs[reflect.Type](t) == Abs[reflect.Type](value.Type()) {
		return structs.ptr(t, value)
	}

	abs := Abs[reflect.Value](value)
	if abs.CanConvert(t) {
		return abs.Convert(t)
	}

	return value
}

func (structs *Struct) Sub(i int, v interface{}) {
	s, ok := structs.Get(i)
	if !ok {
		return
	}

	s.(*Struct).As(&v)
}

func As[T interface{}](v interface{}) T {
	switch v := v.(type) {
	case *Struct:
		var t T

		v.As(&t)
		return t
	case T:
		return v
	default:
		value := Value(v)
		t := reflect.TypeFor[T]()

		var ptr reflect.Value

		switch value.Kind() {
		case reflect.Array, reflect.Slice:
			ptr = as2(value, reflect.New(t).Elem())
		case reflect.Map:
			ptr = as2(value, reflect.MakeMapWithSize(t, value.Len()))
		default:
			var zero T
			return zero
		}

		return ptr.Interface().(T)
	}
}

func as2(src, dst reflect.Value) reflect.Value {
	if s, ok := src.Interface().(*Struct); ok {
		s.As(dst)
		return dst
	}

	src = Abs[reflect.Value](src)
	dst = Abs[reflect.Value](dst)

	switch dst.Type().Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < src.Len(); i++ {
			ptr := as2(src.Index(i), reflect.New(dst.Type().Elem()).Elem())
			dst = reflect.Append(dst, ptr)
		}

		return dst
	case reflect.Map:
		m := src.MapRange()

		for m.Next() {
			k := as2(m.Key(), reflect.New(dst.Type().Key()).Elem())
			v := as2(m.Value(), reflect.New(dst.Type().Elem()).Elem())

			dst.SetMapIndex(k, v)
		}

		return dst
	default:
		return src
	}
}
