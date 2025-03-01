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

package kind

import (
	"reflect"
	"sync"
)

type Data struct {
	Kind    int
	Type    reflect.Type
	Handler Handler
}

type Map struct {
	mkind sync.Map
	mtype sync.Map
	cache sync.Map
}

func (m *Map) Store(kind int, t reflect.Type, handler Handler) {
	data := &Data{
		Kind:    kind,
		Type:    t,
		Handler: handler,
	}

	m.mkind.Store(kind, data)
	m.mtype.Store(t, data)
}

func (m *Map) Load(v interface{}) (int, reflect.Type) {
	switch v.(type) {
	case int:
		kind, ok := m.mkind.Load(v)
		if !ok {
			return 0, nil
		}

		data, ok := kind.(*Data)
		if !ok {
			return 0, nil
		}

		return data.Kind, data.Type
	case reflect.Type:
		if _, ok := m.cache.Load(v); ok {
			return 0, nil
		}

		t, ok := m.mtype.Load(v)
		if !ok {
			m.cache.Store(v, true)
			return 0, nil
		}

		data, ok := t.(*Data)
		if !ok {
			return 0, nil
		}

		return data.Kind, data.Type
	default:
		return 0, nil
	}
}

func (m *Map) Has(t reflect.Type) (int, bool) {
	if _, ok := m.cache.Load(t); ok {
		return 0, false
	}

	data, ok := m.mtype.Load(t)
	if !ok {
		kind := 0
		found := false

		m.mtype.Range(func(rk, rv any) bool {
			if rk.(reflect.Type).Kind() != reflect.Interface {
				return true
			}

			if t.Implements(rk.(reflect.Type)) {
				kind = rv.(*Data).Kind
				found = true

				return false
			}

			return true
		})

		if !found {
			if t.Kind() == reflect.Pointer {
				m.cache.Store(t, true)
				return 0, false
			}

			return m.Has(reflect.PointerTo(t))
		}

		return kind, found
	}

	return data.(*Data).Kind, true
}

func (m *Map) Alias(kind int, t reflect.Type) {
	data, ok := m.mkind.Load(kind)
	if !ok {
		return
	}

	m.mtype.Store(t, data)
}

func (m *Map) Run(v, i interface{}, value reflect.Value) (bool, error) {
	var data *Data

	switch v.(type) {
	case int:
		v, ok := m.mkind.Load(v)
		if !ok {
			return false, nil
		}

		data = v.(*Data)
	case reflect.Type:
		t, ok := m.mtype.Load(v)
		if !ok {
			m.mtype.Range(func(rk, rv any) bool {
				if rk.(reflect.Type).Kind() != reflect.Interface {
					return true
				}

				if value.Type().Implements(rk.(reflect.Type)) {
					t = rv
					return false
				}

				return true
			})

			if t == nil {
				if value.Kind() == reflect.Pointer {
					return false, nil
				}

				ptr := Pointer(value)

				status, err := m.Run(v, i, ptr)
				if status == false || err != nil {
					return status, err
				}

				value.Set(ptr.Elem())
				return status, err
			}
		}

		data = t.(*Data)
	default:
		return false, nil
	}

	switch i := i.(type) {
	case Encoder:
		return true, data.Handler.Encode(i, value)
	case Decoder:
		return true, data.Handler.Decode(i, value)
	default:
		return false, nil
	}
}
