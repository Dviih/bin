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

import "reflect"

type Encoder interface {
	Encode(interface{}) error
}

type Decoder interface {
	Decode(interface{}) error
}

type Handler interface {
	Encode(Encoder, reflect.Value) error
	Decode(Decoder, reflect.Value) error
}

type ed struct {
	encode func(Encoder, reflect.Value) error
	decode func(Decoder, reflect.Value) error
}

func (m *ed) Encode(encoder Encoder, value reflect.Value) error {
	return m.encode(encoder, value)
}

func (m *ed) Decode(decoder Decoder, value reflect.Value) error {
	return m.decode(decoder, value)
}

func NewHandler(encode func(Encoder, reflect.Value) error, decode func(Decoder, reflect.Value) error) Handler {
	return &ed{
		encode: encode,
		decode: decode,
	}
}

// Pointer transforms value T into *T with reflection.
func Pointer(value reflect.Value) reflect.Value {
	ptr := reflect.New(value.Type())
	ptr.Elem().Set(value)

	return ptr
}

// Call tries to call normal value and then pointer.
func Call(value reflect.Value, method string, v ...reflect.Value) []reflect.Value {
	if method == "" {
		return value.Call(v)
	}

	var ptr reflect.Value

	m := value.MethodByName(method)
	if !m.IsValid() {
		ptr = Pointer(value)
		m = ptr.MethodByName(method)
	}

	out := m.Call(v)
	if ptr.Kind() == reflect.Invalid {
		return out
	}

	if value.CanSet() {
		value.Set(ptr.Elem())
	}

	return out
}
