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

type stream struct {
	Data []byte
	read int
}

func (stream *stream) Write(data []byte) (int, error) {
	stream.Data = append(stream.Data, data...)
	return len(data), nil
}

func (stream *stream) Read(data []byte) (int, error) {
	i := copy(data, stream.Data[stream.read:stream.read+len(data)])
	stream.read += len(data)

	return i, nil
}

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

