/*
 *     A tiny binary format
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
	"github.com/Dviih/bin/buffer"
	"testing"
)

func TestEncoderNil(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Nil); err != nil {
		t.Error("failed to encode nil")
	}

	if string(b.Data()) != string(expectedNil) {
		t.Errorf("expected %v, received: %v", expectedNil, b.Data())
	}
}

func TestEncoderBool(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Bool); err != nil {
		t.Error("failed to encode boolean")
	}

	if string(b.Data()) != string(expectedBool) {
		t.Errorf("expected %v, received: %v", expectedBool, b.Data())
	}
}

func TestEncoderInt(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Int); err != nil {
		t.Error("failed to encode int")
	}

	if string(b.Data()) != string(expectedInt) {
		t.Errorf("expected %v, received: %v", expectedInt, b.Data())
	}
}

func TestEncoderUint(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Uint); err != nil {
		t.Error("failed to encode uint")
	}

	if string(b.Data()) != string(expectedUint) {
		t.Errorf("expected %v, received: %v", expectedUint, b.Data())
	}
}

func TestEncoderFloat(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Float); err != nil {
		t.Error("failed to encode float")
	}

	if string(b.Data()) != string(expectedFloat) {
		t.Errorf("expected %v, received: %v", expectedFloat, b.Data())
	}
}

func TestEncoderComplex(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Complex); err != nil {
		t.Error("failed to encode complex")
	}

	if string(b.Data()) != string(expectedComplex) {
		t.Errorf("expected %v, received: %v", expectedComplex, b.Data())
	}
}

func TestEncoderArray(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Array); err != nil {
		t.Error("failed to encode array")
	}

	if string(b.Data()) != string(expectedArray) {
		t.Errorf("expected %v, received: %v", expectedArray, b.Data())
	}
}

func TestEncoderMap(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Map); err != nil {
		t.Error("failed to encode map")
	}

	if string(b.Data()) != string(expectedMap) {
		t.Errorf("expected %v, received: %v", expectedMap, b.Data())
	}
}

func TestEncoderSlice(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Slice); err != nil {
		t.Error("failed to encode slice")
	}

	if string(b.Data()) != string(expectedSlice) {
		t.Errorf("expected %v, received: %v", expectedSlice, b.Data())
	}
}

func TestEncoderString(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(String); err != nil {
		t.Error("failed to encode string")
	}

	if string(b.Data()) != string(expectedString) {
		t.Errorf("expected %v, received: %v", expectedString, b.Data())
	}
}

func TestEncoderStruct(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Struct2); err != nil {
		t.Error("failed to encode struct")
	}

	if string(b.Data()) != string(expectedStruct) {
		t.Errorf("expected %v, received: %v", expectedStruct, b.Data())
	}
}

func TestEncoderStructNumbers(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(StructNumbersValue); err != nil {
		t.Error("failed to encode struct numbers")
	}

	if string(b.Data()) != string(expectedStructNumbers) {
		t.Errorf("expected %v, received: %v", expectedStructNumbers, b.Data())
	}
}

func TestEncoderStructArray(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(StructArrayValue); err != nil {
		t.Error("failed to encode struct array")
	}

	if string(b.Data()) != string(expectedStructArray) {
		t.Errorf("expected %v, received: %v", expectedStructArray, b.Data())
	}
}

func TestEncoderStructMap(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(StructMapValue); err != nil {
		t.Error("failed to encode struct map")
	}

	if string(b.Data()) != string(expectedStructMap) {
		t.Errorf("expected %v, received: %v", expectedStructMap, b.Data())
	}
}

func TestEncoderStructAll(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(StructAllValue); err != nil {
		t.Error("failed to encode struct all")
	}

	if string(b.Data()) != string(expectedStructAll) {
		t.Errorf("expected %v, received: %v", expectedStructAll, b.Data())
	}
}

func TestEncoderInterfaceNil(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Interface(nil)); err != nil {
		t.Errorf("failed to encode interface nil")
	}

	if string(b.Data()) != string(expectedInterfaceNil) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceNil, b.Data())
	}
}

func TestEncoderInterfaceBool(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Interface(Bool)); err != nil {
		t.Errorf("failed to encode interface boolean")
	}

	if string(b.Data()) != string(expectedInterfaceBool) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceBool, b.Data())
	}
}

func TestEncoderInterfaceArray(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Interface(Array)); err != nil {
		t.Errorf("failed to encode interface array")
	}

	if string(b.Data()) != string(expectedInterfaceArray) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceArray, b.Data())
	}
}

func TestEncoderInterfaceMap(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Interface(Map)); err != nil {
		t.Errorf("failed to encode interface map")
	}

	if string(b.Data()) != string(expectedInterfaceMap) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceMap, b.Data())
	}
}

// Key and Value are `interface{}`.
func TestEncoderInterfaceMap2(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Interface(Map2)); err != nil {
		t.Errorf("failed to encode interface map2")
	}

	if string(b.Data()) != string(expectedInterfaceMap2) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceMap2, b.Data())
	}
}

// Only Key is `interface{}`.
func TestEncoderInterfaceMap3(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Interface(Map3)); err != nil {
		t.Errorf("failed to encode interface map3")
	}

	if string(b.Data()) != string(expectedInterfaceMap3) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceMap3, b.Data())
	}
}

// Only Value is `interface{}`.
func TestEncoderInterfaceMap4(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Interface(Map4)); err != nil {
		t.Errorf("failed to encode interface map4")
	}

	if string(b.Data()) != string(expectedInterfaceMap4) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceMap4, b.Data())
	}
}

func TestEncoderInterfaceSlice(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Interface(Slice)); err != nil {
		t.Errorf("failed to encode interface slice")
	}

	if string(b.Data()) != string(expectedInterfaceSlice) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceSlice, b.Data())
	}
}

func TestEncoderInterfaceSlice2(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Interface(Slice2)); err != nil {
		t.Errorf("failed to encode interface slice2")
	}

	if string(b.Data()) != string(expectedInterfaceSlice2) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceSlice2, b.Data())
	}
}

func TestEncoderInterfaceString(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Interface(String)); err != nil {
		t.Errorf("failed to encode interface string")
	}

	if string(b.Data()) != string(expectedInterfaceString) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceString, b.Data())
	}
}

func TestEncoderInterfaceStruct(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Interface(Struct2)); err != nil {
		t.Errorf("failed to encode interface struct")
	}

	if string(b.Data()) != string(expectedInterfaceStruct) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceStruct, b.Data())
	}
}

func TestEncoderInterfaceStructNumbers(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Interface(StructNumbersValue)); err != nil {
		t.Errorf("failed to encode interface struct numbers")
	}

	if string(b.Data()) != string(expectedInterfaceStructNumbers) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceStructNumbers, b.Data())
	}
}

func TestEncoderInterfaceStructArray(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Interface(StructArrayValue)); err != nil {
		t.Errorf("failed to encode interface struct array")
	}

	if string(b.Data()) != string(expectedInterfaceStructArray) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceStructArray, b.Data())
	}
}
func TestEncoderInterfaceStructMap(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Interface(StructMapValue)); err != nil {
		t.Errorf("failed to encode interface struct map")
	}

	if string(b.Data()) != string(expectedInterfaceStructMap) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceStructMap, b.Data())
	}
}

func TestEncoderInterfaceStructAll(t *testing.T) {
	t.Parallel()

	b := buffer.New()
	encoder := NewEncoder(b)

	if err := encoder.Encode(Interface(StructAllValue)); err != nil {
		t.Errorf("failed to encode interface struct all")
	}

	if string(b.Data()) != string(expectedInterfaceStructAll) {
		t.Errorf("expected: %v, received: %v", expectedInterfaceStructAll, b.Data())
	}
}
