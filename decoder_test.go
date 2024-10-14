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

func TestDecoderArray(t *testing.T) {
	t.Parallel()

	s := &stream{
		Data: expectedArray,
	}

	decoder := NewDecoder(s)

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

	s := &stream{
		Data: expectedMap,
	}

	decoder := NewDecoder(s)

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

	s := &stream{
		Data: expectedSlice,
	}

	decoder := NewDecoder(s)

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

	s := &stream{
		Data: expectedString,
	}

	decoder := NewDecoder(s)

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

	s := &stream{
		Data: expectedStruct,
	}

	decoder := NewDecoder(s)

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

	s := &stream{
		Data: expectedStructNumbers,
	}

	decoder := NewDecoder(s)

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

	s := &stream{
		Data: expectedStructArray,
	}

	decoder := NewDecoder(s)

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

	s := &stream{
		Data: expectedStructMap,
	}

	decoder := NewDecoder(s)

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

	s := &stream{
		Data: expectedStructAll,
	}

	decoder := NewDecoder(s)

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

	s := &stream{
		Data: expectedNil,
	}

	decoder := NewDecoder(s)

	ptr := reflect.New(reflect.TypeFor[interface{}]()).Elem()
	if err := decoder.Decode(ptr); err != nil {
		t.Error("failed to decode nil")
	}

	if !ptr.IsNil() {
		t.Errorf("expected %v, received: %v", Nil, ptr)
	}
}

