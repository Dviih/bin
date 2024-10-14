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

func TestEncoderString(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(String); err != nil {
		t.Error("failed to encode string")
	}

	if string(s.Data) != string(expectedString) {
		t.Errorf("expected %v, received: %v", expectedString, s.Data)
	}
}

func TestEncoderStruct(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(Struct2); err != nil {
		t.Error("failed to encode struct")
	}

	if string(s.Data) != string(expectedStruct) {
		t.Errorf("expected %v, received: %v", expectedStruct, s.Data)
	}
}

func TestEncoderStructNumbers(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(StructNumbersValue); err != nil {
		t.Error("failed to encode struct numbers")
	}

	if string(s.Data) != string(expectedStructNumbers) {
		t.Errorf("expected %v, received: %v", expectedStructNumbers, s.Data)
	}
}

func TestEncoderStructArray(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(StructArrayValue); err != nil {
		t.Error("failed to encode struct array")
	}

	if string(s.Data) != string(expectedStructArray) {
		t.Errorf("expected %v, received: %v", expectedStructArray, s.Data)
	}
}

func TestEncoderStructMap(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(StructMapValue); err != nil {
		t.Error("failed to encode struct map")
	}

	if string(s.Data) != string(expectedStructMap) {
		t.Errorf("expected %v, received: %v", expectedStructMap, s.Data)
	}
}

func TestEncoderStructAll(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(StructAllValue); err != nil {
		t.Error("failed to encode struct all")
	}

	if string(s.Data) != string(expectedStructAll) {
		t.Errorf("expected %v, received: %v", expectedStructAll, s.Data)
	}
}

func TestEncoderInterfaceNil(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(Interface(nil)); err != nil {
		t.Errorf("failed to encode interface nil")
	}

	if string(s.Data) != string(expectedInterfaceNil) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceNil, s.Data)
	}
}

func TestEncoderInterfaceBool(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(Interface(Bool)); err != nil {
		t.Errorf("failed to encode interface boolean")
	}

	if string(s.Data) != string(expectedInterfaceBool) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceBool, s.Data)
	}
}

