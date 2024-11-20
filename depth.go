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
	"slices"
)

func depth(value reflect.Value) (reflect.Type, int, bool, []int) {
	i := 0
	t := value.Type()

	var di []int
	mixed := isMixed(t)

	for {
		switch t.Kind() {
		case reflect.Array:
			di = append(di, t.Len())

			i++
			t = t.Elem()
		case reflect.Slice:
			if mixed {
				di = append(di, 0)
			} else {
				di = append(di, value.Len())
			}

			i++
			value = value.Index(0)
			t = t.Elem()
		default:
			return t, i, mixed, di
		}
	}
}

func isMixed(t reflect.Type) bool {
	pt := t

	for {
		switch t.Elem().Kind() {
		case reflect.Array, reflect.Slice:
			if t.Elem().Kind() == pt.Kind() {
				pt = t
				t = t.Elem()
				continue
			}
			return true
		default:
			return false
		}
	}
}

