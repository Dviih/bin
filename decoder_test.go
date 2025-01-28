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
	"github.com/Dviih/bin/buffer"
	"reflect"
	"testing"
)

func TestDecoderNil(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedNil))

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

	decoder := NewDecoder(buffer.From(expectedBool))

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

	decoder := NewDecoder(buffer.From(expectedInt))

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

	decoder := NewDecoder(buffer.From(expectedUint))

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

	decoder := NewDecoder(buffer.From(expectedFloat))

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

	decoder := NewDecoder(buffer.From(expectedComplex))

	var c complex128
	if err := decoder.Decode(&c); err != nil {
		t.Error("failed to decode complex")
	}

	if c != Complex {
		t.Errorf("expected %v, received: %v", Complex, c)
	}
}

func TestDecoderArray(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedArray))

	var array [3]uint64
	if err := decoder.Decode(&array); err != nil {
		t.Error("failed to decode array")
	}

	if array != Array {
		t.Errorf("expected %v, received: %v", Array, array)
	}
}

func TestDecoderMap(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedMap))

	var m map[byte]int
	if err := decoder.Decode(&m); err != nil {
		t.Error("failed to decode map")
	}

	if !reflect.DeepEqual(m, Map) {
		t.Errorf("expected %v, received: %v", Map, m)
	}
}

func TestDecoderSlice(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedSlice))

	var slice []int
	if err := decoder.Decode(&slice); err != nil {
		t.Error("failed to decode slice")
	}

	if !reflect.DeepEqual(slice, Slice) {
		t.Errorf("expected %v, received: %v", Slice, slice)
	}
}

func TestDecoderString(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedString))

	var _string string
	if err := decoder.Decode(&_string); err != nil {
		t.Error("failed to decode string")
	}

	if _string != String {
		t.Errorf("expected %v, received: %v", String, _string)
	}
}

func TestDecoderStruct(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedStruct))

	var struct2 *Struct1
	if err := decoder.Decode(&struct2); err != nil {
		t.Error("failed to decode struct")
	}

	if !reflect.DeepEqual(struct2, Struct2) {
		t.Errorf("expected %v, received: %v", Struct2, struct2)
	}
}

func TestDecoderStructNumbers(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedStructNumbers))

	var structNumbers *StructNumbers
	if err := decoder.Decode(&structNumbers); err != nil {
		t.Error("failed to decode struct numbers")
	}

	if !reflect.DeepEqual(structNumbers, StructNumbersValue) {
		t.Errorf("expected %v, received: %v", StructNumbersValue, structNumbers)
	}
}

func TestDecoderStructArray(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedStructArray))

	var structArray *StructArray
	if err := decoder.Decode(&structArray); err != nil {
		t.Error("failed to decode struct array")
	}

	if !reflect.DeepEqual(structArray, StructArrayValue) {
		t.Errorf("expected %v, received: %v", StructArrayValue, structArray)
	}
}

func TestDecoderStructMap(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedStructMap))

	var structMap *StructMap
	if err := decoder.Decode(&structMap); err != nil {
		t.Error("failed to decode struct map")
	}

	if !reflect.DeepEqual(structMap, StructMapValue) {
		t.Errorf("expected %v, received: %v", StructMapValue, structMap)
	}
}

func TestDecoderStructAll(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedStructAll))

	var structAll *StructAll
	if err := decoder.Decode(&structAll); err != nil {
		t.Error("failed to decode struct all")
	}

	if !reflect.DeepEqual(structAll, StructAllValue) {
		t.Errorf("expected %v, received: %v", StructAllValue, structAll)
	}
}

func TestDecoderInterfaceNil(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedInterfaceNil))

	ptr := reflect.New(reflect.TypeFor[interface{}]()).Elem()
	if err := decoder.Decode(ptr); err != nil {
		t.Error("failed to decode nil")
	}

	if !ptr.IsNil() {
		t.Errorf("expected %v, received: %v", Nil, ptr)
	}
}

func TestDecoderInterfaceBool(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedInterfaceBool))

	var i interface{}
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode bool")
	}

	if i.(bool) != Bool {
		t.Errorf("expected %v, received: %v", Bool, i)
	}
}

func TestDecoderInterfaceArray(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedInterfaceArray))

	var i interface{}
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode array")
	}

	if !reflect.DeepEqual(i, Array) {
		t.Errorf("expected %v, received: %v", Array, i)
	}
}

func TestDecoderInterfaceMap(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedInterfaceMap))

	var i interface{}
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode map")
	}

	if !reflect.DeepEqual(i, Map) {
		t.Errorf("expected %v, received: %v", Map, i)
	}
}

func TestDecoderInterfaceMap2(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedInterfaceMap2))

	var i interface{}
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode map2")
	}

	if !reflect.DeepEqual(i, Map2) {
		t.Errorf("expected %v, received: %v", Map2, i)
	}
}

func TestDecoderInterfaceMap3(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedInterfaceMap3))

	var i interface{}
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode map3")
	}

	if !reflect.DeepEqual(i, Map3) {
		t.Errorf("expected %v, received: %v", Map3, i)
	}
}

func TestDecoderInterfaceMap4(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedInterfaceMap4))

	var i interface{}
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode map4")
	}

	if !reflect.DeepEqual(i, Map4) {
		t.Errorf("expected %v, received: %v", Map4, i)
	}
}

func TestDecoderInterfaceSlice(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedInterfaceSlice))

	var i interface{}
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode slice")
	}

	if !reflect.DeepEqual(i, Slice) {
		t.Errorf("expected %v, received: %v", Slice, i)
	}
}

func TestDecoderInterfaceSlice2(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedInterfaceSlice2))

	var i interface{}
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode slice2")
	}

	if !reflect.DeepEqual(i, Slice2) {
		t.Errorf("expected %v, received: %v", Slice2, i)
	}
}

func TestDecoderInterfaceString(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedInterfaceString))

	var i interface{}
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode string")
	}

	if i.(string) != String {
		t.Errorf("expected %v, received: %v", String, i)
	}
}

func TestDecoderInterfaceStruct(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedInterfaceStruct))

	var i interface{}
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode struct")
	}

	st := As[*Struct1](i)

	if !reflect.DeepEqual(st, Struct2) {
		t.Errorf("expected %v, received: %v", Struct2, st)
	}
}

func TestDecoderInterfaceStructNumbers(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedInterfaceStructNumbers))

	var i interface{}
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode struct numbers")
	}

	var st *StructNumbers
	i.(*Struct).As(&st)

	if !reflect.DeepEqual(st, StructNumbersValue) {
		t.Errorf("expected %v, received: %v", StructNumbersValue, st)
	}
}

func TestDecoderInterfaceStructArray(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedInterfaceStructArray))

	var i interface{}
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode struct array")
	}

	var st *StructArray
	i.(*Struct).As(&st)

	if !reflect.DeepEqual(st, StructArrayValue) {
		t.Errorf("expected %v, received: %v", StructArrayValue, st)
	}
}

func TestDecoderInterfaceStructMap(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedInterfaceStructMap))

	var i interface{}
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode struct map")
	}

	var st *StructMap
	i.(*Struct).As(&st)

	if !reflect.DeepEqual(st, StructMapValue) {
		t.Errorf("expected %v, received: %v", StructMapValue, st)
	}
}

func TestDecoderInterfaceStructAll(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(buffer.From(expectedInterfaceStructAll))

	var i interface{}
	if err := decoder.Decode(&i); err != nil {
		t.Error("failed to decode struct all")
	}

	st := As[*StructAll](i)

	if !reflect.DeepEqual(st, StructAllValue) {
		t.Errorf("expected %v, received: %v", StructAllValue, st)
	}
}
