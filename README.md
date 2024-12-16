# Bin

## A tiny binary format
##### This project is licensed under [AGPL](https://github.com/Dviih/bin/blob/main/LICENSE).

---

## Encoder
### This implements encoding and takes an `io.Writer`.

## Decoder
### This implements decoding and takes an `io.Reader`.

## Struct
### A map where it can translate into a struct with the help of the tag.

---

# Usage

## Encoder
- `Encode` - Takes an `interface{}` and writes to `io.Writer`, returns an error if value is invalid.
- `getType` - Takes the type from a `reflect.Value`, only used for interfaced values.
- `structs` - Takes a `reflect.Value` and a boolean if kind is required to write into `io.Writer`, returns an error if value is either invalid or tag is not a number.

## Decoder
- `Decode` - Takes an `interface{}` and decodes from `io.Reader`, returs an error if value is invalid or value is not settable or `io.Reader` read an invalid VarUint.
- `getType` - Decodes the type and heads towards decoding it, only used for interfaced values.
- `ReadByte` - Returns a byte and an `io.EOF` if `io.Read` is done.
- `structs` - Takes a `reflect.Value` and decodes each struct field, might return an error as the same for `Decode`.

## Struct

##
- `Map` - Returns a map representing the struct.
- `Get` - Returns the key and a status.
- `As` - Takes an `interface{}` and sets what the `interface{}` has, it will do nothing if the interface is not a struct.
- `Sub` - Takes a tag and an `interface{}` and sets the interface with the value from the tag.
- `maps` - Takes a map and ranges through its values returning a readable map.
- `fields` - Takes a struct and map its fields by their tag, returns `map[tag]field`
- `ranges` - Takes `fields` map and check the values if possible to set, converting if required, if not continue.
- `convert` - Safe conversion for `reflect.Value`.
- `ptr` - Matches the `reflect.Value` of the field.

---

## Utilities
#### This section is used to describe utilities used in the project, some are available, some are not.

## General utilities

- `Value` - Takes an `interface{}` and returns `reflect.Value`, if interface is already return it back, else get the absolute value of the `interface{}`.
- `Zero` - Initializes a `reflect.Value`.
- `Abs[T]` - Gets absolute `reflect.Type` or `reflect.Value`.

## Marshaling and Unmarshaling utilities
- `Marshal` - Takes `interface{}` and returns bytes, returns error as the same as Encoder.
- `Unmarshal[T]` - Takes `[]byte` and decodes into T, returns error as the same as Decoder.
- `UnmarshalAs[T]` - Combines `Unmarshal[T]` and `As[T]` calls.

## `interface{}` utilities.

- `Interface` - Takes an `interface{}` and returns `reflect.Value`.
- `interfaces` - Takes a `reflect.Value` and switches to `interface{}` whatever is needed, returns itself.
- `As[T]` - Takes an `interface{}` and tries to decode into T, if so returns it.
- `as2` - Takes a `reflect.Type` and `interface{}` being the head behind `As[T]`.
- `KeyElem` - Takes a `reflect.Value` expecting arrays, slices and maps and returns its key and element types.

## VarInt utilities

- `Integer` - Integer is an interface with all integer types of Go also allowing other types with integers underlying it.
- `VarIntIn[Integer]` - Takes int and uint ranges and an `io.Writer` and returns bytes, error if `io.Writer` is done.
- `VarIntOut[Integer]` - Must disclosure the type and takes an `io.ByteReader`, returns the number and an error if `io.Reader` of Decoder is done.

## Depth utilities

- `depth` - Takes a `reflect.Value` kind must be either `reflect.Array` or `reflect.Slice` and calculates depth, mixed state and depth sizes.
- `fromDepth` - Builds a type from depth.
- `isMixed` ` Check if all arrays are an arrays or slices are all slices, if so returns true.

---

# Example

```go
package main

import (
	"fmt"
	"github.com/Dviih/bin"
	"reflect"
)

type Example struct {
	Twenty  int    `bin:"20"`
	Fifth   []int  `bin:"50"`
	Hundred string `bin:"100"`
}

func main() {
	example := &Example{
		Twenty:  20,
		Fifth:   []int{5, 0},
		Hundred: "A hundred",
	}

	data, err := bin.Marshal(example)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)

	example2, err := bin.Unmarshal[*Example](data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", example2)
	fmt.Println("Equal", reflect.DeepEqual(example, example2))
}
```
###### More can be found at `bin_test.go`, `encoder_test.go` and `decoder_test.go`

---

#### Made for Gophers by @Dviih