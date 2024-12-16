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
	"reflect"
	"testing"
)

type Struct1 struct {
	FieldOne string `bin:"100"`
	FieldTwo uint64 `bin:"200"`
}

type StructNumbers struct {
	Int        int        `bin:"10"`
	Int8       int8       `bin:"20"`
	Int16      int16      `bin:"30"`
	Int32      int32      `bin:"40"`
	Int64      int64      `bin:"50"`
	Uint       uint       `bin:"60"`
	Uint8      uint8      `bin:"70"`
	Uint16     uint16     `bin:"80"`
	Uint32     uint32     `bin:"90"`
	Uint64     uint64     `bin:"100"`
	Float32    float32    `bin:"110"`
	Float64    float64    `bin:"120"`
	Complex64  complex64  `bin:"130"`
	Complex128 complex128 `bin:"140"`
}

type StructArray struct {
	Numbers []int         `bin:"10"`
	Stuff   []interface{} `bin:"20"`
}

type StructMap struct {
	Data  map[byte]uint64             `bin:"10"`
	Stuff map[interface{}]interface{} `bin:"20"`
}

type StructAll struct {
	One   *Struct1
	Two   *StructNumbers
	Three *StructArray
	Four  *StructMap
}

var (
	Nil     interface{}
	Bool    = true
	Int     = 768
	Uint    = uint(10240)
	Float   = 42.69
	Complex = complex(69, 42)
	Array   = [3]uint64{1, 256, 1024}
	Map     = map[byte]int{
		16: 1024,
	}
	Map2 = map[interface{}]interface{}{
		"string": 10,
	}
	Map3 = map[interface{}]string{
		20: "twenty",
	}
	Map4 = map[string]interface{}{
		"fifth": 50,
	}
	Slice   = []int{24, 69, 128, 512}
	Slice2  = []interface{}{"twenty", 50, "hundreds"}
	String  = "Hello, World!"
	Struct2 = &Struct1{
		FieldOne: "one",
		FieldTwo: 2,
	}
	StructNumbersValue = &StructNumbers{
		Int:        1,
		Int8:       2,
		Int16:      4,
		Int32:      8,
		Int64:      16,
		Uint:       32,
		Uint8:      64,
		Uint16:     128,
		Uint32:     256,
		Uint64:     512,
		Float32:    10.24,
		Float64:    20.48,
		Complex64:  complex(40, 96),
		Complex128: complex(81, 92),
	}
	StructArrayValue = &StructArray{
		Numbers: []int{3, 9, 27, 81},
		Stuff:   []interface{}{"Hello", 13, "World", "!"},
	}
	StructMapValue = &StructMap{
		Data: map[byte]uint64{
			10: 1024,
		},
		Stuff: map[interface{}]interface{}{
			81: "nine",
		},
	}
	StructAllValue = &StructAll{
		One:   Struct2,
		Two:   StructNumbersValue,
		Three: StructArrayValue,
		Four:  StructMapValue,
	}

	expectedNil           = []byte{0}
	expectedBool          = []byte{255}
	expectedInt           = []byte{128, 6}
	expectedUint          = []byte{128, 80}
	expectedFloat         = []byte{184, 189, 148, 220, 158, 138, 214, 162, 64}
	expectedComplex       = []byte{128, 128, 128, 128, 128, 128, 208, 168, 64, 128, 128, 128, 128, 128, 128, 192, 162, 64}
	expectedArray         = []byte{1, 128, 2, 128, 8}
	expectedMap           = []byte{1, 16, 128, 8}
	expectedSlice         = []byte{4, 24, 69, 128, 1, 128, 4}
	expectedString        = []byte{13, 72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	expectedStruct        = []byte{100, 3, 111, 110, 101, 200, 1, 2}
	expectedStructNumbers = []byte{10, 1, 20, 2, 30, 4, 40, 8, 50, 16, 60, 32, 70, 64, 80, 128, 1, 90, 128, 2, 100, 128, 4, 110, 138, 174, 143, 137, 4, 120, 251, 168, 184, 189, 148, 220, 158, 154, 64, 130, 1, 128, 128, 128, 145, 4, 128, 128, 128, 150, 4, 140, 1, 128, 128, 128, 128, 128, 128, 144, 170, 64, 128, 128, 128, 128, 128, 128, 192, 171, 64}
	expectedStructArray   = []byte{10, 4, 3, 9, 27, 81, 20, 4, 24, 5, 72, 101, 108, 108, 111, 2, 13, 24, 5, 87, 111, 114, 108, 100, 24, 1, 33}
	expectedStructMap     = []byte{10, 1, 10, 128, 8, 20, 1, 2, 81, 24, 4, 110, 105, 110, 101}
	expectedStructAll     = []byte{1, 100, 3, 111, 110, 101, 200, 1, 2, 2, 10, 1, 20, 2, 30, 4, 40, 8, 50, 16, 60, 32, 70, 64, 80, 128, 1, 90, 128, 2, 100, 128, 4, 110, 138, 174, 143, 137, 4, 120, 251, 168, 184, 189, 148, 220, 158, 154, 64, 130, 1, 128, 128, 128, 145, 4, 128, 128, 128, 150, 4, 140, 1, 128, 128, 128, 128, 128, 128, 144, 170, 64, 128, 128, 128, 128, 128, 128, 192, 171, 64, 3, 10, 4, 3, 9, 27, 81, 20, 4, 24, 5, 72, 101, 108, 108, 111, 2, 13, 24, 5, 87, 111, 114, 108, 100, 24, 1, 33, 4, 10, 1, 10, 128, 8, 20, 1, 2, 81, 24, 4, 110, 105, 110, 101}

	expectedInterfaceNil           = append([]byte{byte(reflect.Invalid)}, expectedNil...)
	expectedInterfaceBool          = append([]byte{byte(reflect.Bool)}, expectedBool...)
	expectedInterfaceArray         = append([]byte{byte(reflect.Array), 1, 0, 3, byte(reflect.Uint64)}, expectedArray...)
	expectedInterfaceMap           = append([]byte{byte(reflect.Map), byte(reflect.Uint8), byte(reflect.Int)}, expectedMap...)
	expectedInterfaceMap2          = []byte{21, 20, 20, 1, 24, 6, 115, 116, 114, 105, 110, 103, 2, 10}
	expectedInterfaceMap3          = []byte{21, 20, 24, 1, 2, 20, 6, 116, 119, 101, 110, 116, 121}
	expectedInterfaceMap4          = []byte{21, 24, 20, 1, 5, 102, 105, 102, 116, 104, 2, 50}
	expectedInterfaceSlice         = append([]byte{byte(reflect.Slice), 1, 0, byte(reflect.Int)}, expectedSlice...)
	expectedInterfaceSlice2        = []byte{23, 1, 0, 20, 3, 24, 6, 116, 119, 101, 110, 116, 121, 2, 50, 24, 8, 104, 117, 110, 100, 114, 101, 100, 115}
	expectedInterfaceString        = append([]byte{byte(reflect.String)}, expectedString...)
	expectedInterfaceStruct        = []byte{25, 2, 100, 24, 3, 111, 110, 101, 200, 1, 11, 2}
	expectedInterfaceStructNumbers = []byte{25, 14, 10, 2, 1, 20, 3, 2, 30, 4, 4, 40, 5, 8, 50, 6, 16, 60, 7, 32, 70, 8, 64, 80, 9, 128, 1, 90, 10, 128, 2, 100, 11, 128, 4, 110, 13, 138, 174, 143, 137, 4, 120, 14, 251, 168, 184, 189, 148, 220, 158, 154, 64, 130, 1, 15, 128, 128, 128, 145, 4, 128, 128, 128, 150, 4, 140, 1, 16, 128, 128, 128, 128, 128, 128, 144, 170, 64, 128, 128, 128, 128, 128, 128, 192, 171, 64}
	expectedInterfaceStructArray   = []byte{25, 2, 10, 23, 1, 0, 2, 4, 3, 9, 27, 81, 20, 23, 1, 0, 20, 4, 24, 5, 72, 101, 108, 108, 111, 2, 13, 24, 5, 87, 111, 114, 108, 100, 24, 1, 33}
	expectedInterfaceStructMap     = []byte{25, 2, 10, 21, 8, 11, 1, 10, 128, 8, 20, 21, 20, 20, 1, 2, 81, 24, 4, 110, 105, 110, 101}
	expectedInterfaceStructAll     = []byte{25, 4, 1, 25, 2, 100, 24, 3, 111, 110, 101, 200, 1, 11, 2, 2, 25, 14, 10, 2, 1, 20, 3, 2, 30, 4, 4, 40, 5, 8, 50, 6, 16, 60, 7, 32, 70, 8, 64, 80, 9, 128, 1, 90, 10, 128, 2, 100, 11, 128, 4, 110, 13, 138, 174, 143, 137, 4, 120, 14, 251, 168, 184, 189, 148, 220, 158, 154, 64, 130, 1, 15, 128, 128, 128, 145, 4, 128, 128, 128, 150, 4, 140, 1, 16, 128, 128, 128, 128, 128, 128, 144, 170, 64, 128, 128, 128, 128, 128, 128, 192, 171, 64, 3, 25, 2, 10, 23, 1, 0, 2, 4, 3, 9, 27, 81, 20, 23, 1, 0, 20, 4, 24, 5, 72, 101, 108, 108, 111, 2, 13, 24, 5, 87, 111, 114, 108, 100, 24, 1, 33, 4, 25, 2, 10, 21, 8, 11, 1, 10, 128, 8, 20, 21, 20, 20, 1, 2, 81, 24, 4, 110, 105, 110, 101}
)

func TestMarshal(t *testing.T) {
	data, err := Marshal(StructNumbersValue)
	if err != nil {
		t.Error("failed to marshal")
	}

	if string(data) != string(expectedStructNumbers) {
		t.Errorf("expected %v, received: %v", expectedStructNumbers, data)
		return
	}
}

func TestUnmarshal(t *testing.T) {
	st, err := Unmarshal[*StructNumbers](expectedStructNumbers)
	if err != nil {
		t.Error("failed to unmarshal")
	}

	if !reflect.DeepEqual(st, StructNumbersValue) {
		t.Errorf("expected %v, received: %v", StructNumbersValue, st)
		return
	}
}

func BenchmarkEncode(b *testing.B) {
	data, err := Marshal(StructAllValue)
	if err != nil {
		b.Error("failed to encode (bench)")
	}

	b.StopTimer()

	if string(data) != string(expectedStructAll) {
		b.Error("not equal (bench)")
	}
}

func BenchmarkDecode(b *testing.B) {
	sa, err := Unmarshal[*StructAll](expectedStructAll)
	if err != nil {
		b.Errorf("failed to unmarshalas (bench): %v", err)
	}

	b.StopTimer()

	if !reflect.DeepEqual(sa, StructAllValue) {
		b.Error("not equal (bench)")
	}
}
