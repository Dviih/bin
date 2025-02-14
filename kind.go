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
	"sync"
)

type Handler interface {
	BinOutput(*Encoder, reflect.Value)
	BinInput(*Decoder, reflect.Value)
}

type kindData struct {
	Kind    int
	Type    reflect.Type
	Handler Handler
}

type kindMap struct {
	// mkind map[int]kindData
	mkind sync.Map

	// mtype map[reflect.Type]kindData
	mtype sync.Map
}

func (m *kindMap) Store(kind int, t reflect.Type, handler Handler) {
	data := &kindData{
		Kind:    kind,
		Type:    t,
		Handler: handler,
	}

	go m.mkind.Store(kind, data)
	go m.mtype.Store(t, data)
}

func (m *kindMap) Get(v interface{}) (int, reflect.Type) {
	switch v.(type) {
	case int:
		kind, ok := m.mkind.Load(v)
		if !ok {
			return 0, nil
		}

		data, ok := kind.(*kindData)
		if !ok {
			return 0, nil
		}
		return data.Kind, data.Type
	case reflect.Type:
		t, ok := m.mtype.Load(v)
		if !ok {
			return 0, nil
		}

		data, ok := t.(*kindData)
		if !ok {
			return 0, nil
		}

		return data.Kind, data.Type
	default:
		return 0, nil
	}
}

