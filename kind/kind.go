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

