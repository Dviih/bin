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
		t, ok := m.mtype.Load(v)
		if !ok {
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
		v, ok := m.mtype.Load(v)
		if !ok {
			return false, nil
		}

		data = v.(*Data)
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
