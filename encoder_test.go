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
	"testing"
)

func TestEncoderNil(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(Nil); err != nil {
		t.Error("failed to encode nil")
	}

	if string(s.Data) != string(expectedNil) {
		t.Errorf("expected %v, received: %v", expectedNil, s.Data)
	}
}

func TestEncoderBool(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(Bool); err != nil {
		t.Error("failed to encode boolean")
	}

	if string(s.Data) != string(expectedBool) {
		t.Errorf("expected %v, received: %v", expectedBool, s.Data)
	}
}

func TestEncoderInt(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(Int); err != nil {
		t.Error("failed to encode int")
	}

	if string(s.Data) != string(expectedInt) {
		t.Errorf("expected %v, received: %v", expectedInt, s.Data)
	}
}

func TestEncoderUint(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(Uint); err != nil {
		t.Error("failed to encode uint")
	}

	if string(s.Data) != string(expectedUint) {
		t.Errorf("expected %v, received: %v", expectedUint, s.Data)
	}
}

func TestEncoderFloat(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(Float); err != nil {
		t.Error("failed to encode float")
	}

	if string(s.Data) != string(expectedFloat) {
		t.Errorf("expected %v, received: %v", expectedFloat, s.Data)
	}
}

func TestEncoderComplex(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(Complex); err != nil {
		t.Error("failed to encode complex")
	}

	if string(s.Data) != string(expectedComplex) {
		t.Errorf("expected %v, received: %v", expectedComplex, s.Data)
	}
}

func TestEncoderArray(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(Array); err != nil {
		t.Error("failed to encode array")
	}

	if string(s.Data) != string(expectedArray) {
		t.Errorf("expected %v, received: %v", expectedArray, s.Data)
	}
}

func TestEncoderMap(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(Map); err != nil {
		t.Error("failed to encode map")
	}

	if string(s.Data) != string(expectedMap) {
		t.Errorf("expected %v, received: %v", expectedMap, s.Data)
	}
}

func TestEncoderSlice(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(Slice); err != nil {
		t.Error("failed to encode slice")
	}

	if string(s.Data) != string(expectedSlice) {
		t.Errorf("expected %v, received: %v", expectedSlice, s.Data)
	}
}

