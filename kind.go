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
	"encoding"
	"github.com/Dviih/bin/kind"
	"reflect"
)

var mkind = &kind.Map{}

func Register[T interface{}](n int, handler kind.Handler) {
	if n < 128 {
		panic("invalid kind range")
	}

	register(n, Abs[reflect.Type](reflect.TypeFor[T]()), handler)
}

func register(n int, t reflect.Type, handler kind.Handler) {
	mkind.Store(n, t, handler)
}

func Alias[T interface{}](n int) {
	if n < 128 {
		panic("invalid kind range")
	}

	mkind.Alias(n, Abs[reflect.Type](reflect.TypeFor[T]()))
}


