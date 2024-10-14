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
	"testing"
)

func TestDecoderNil(t *testing.T) {
	t.Parallel()

	s := &stream{
		Data: expectedNil,
	}

	decoder := NewDecoder(s)

	var i interface{}
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode nil")
	}

	if i != Nil {
		t.Errorf("expected %v, received: %v", Nil, i)
	}
}

func TestDecoderBool(t *testing.T) {
	t.Parallel()

	s := &stream{
		Data: expectedBool,
	}

	decoder := NewDecoder(s)

	var b bool
	if err := decoder.Decode(&b); err != nil {
		t.Error("failed to decode bool")
	}

	if b != Bool {
		t.Errorf("expected %v, received: %v", Bool, b)
	}
}

func TestDecoderInt(t *testing.T) {
	t.Parallel()

	s := &stream{
		Data: expectedInt,
	}

	decoder := NewDecoder(s)

	var i int
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode int")
	}

	if i != Int {
		t.Errorf("expected %v, received: %v", Int, i)
	}
}

func TestDecoderUint(t *testing.T) {
	t.Parallel()

	s := &stream{
		Data: expectedUint,
	}

	decoder := NewDecoder(s)

	var u uint
	if err := decoder.Decode(&u); err != nil {
		t.Error("failed to decode uint")
	}

	if u != Uint {
		t.Errorf("expected %v, received: %v", Uint, u)
	}
}

func TestDecoderFloat(t *testing.T) {
	t.Parallel()

	s := &stream{
		Data: expectedFloat,
	}

	decoder := NewDecoder(s)

	var f float64
	if err := decoder.Decode(&f); err != nil {
		t.Error("failed to decode float")
	}

	if f != Float {
		t.Errorf("expected %v, received: %v", Float, f)
	}
}

func TestDecoderComplex(t *testing.T) {
	t.Parallel()

	s := &stream{
		Data: expectedComplex,
	}

	decoder := NewDecoder(s)

	var c complex128
	if err := decoder.Decode(&c); err != nil {
		t.Error("failed to decode complex")
	}

	if c != Complex {
		t.Errorf("expected %v, received: %v", Complex, c)
	}
}

