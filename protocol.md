# Bin Protocol
### This file describes how bin protocol should be implemented.

---

## Reserved Spaces
### Bin uses the types as `reflect.Kind` does.
- 0-64 reserved for `reflect.Kind`.
- 65-96 reserved for interface implementations.
- 97-127 reserved for custom types handling, can extend to 2^64-1.

## Extensions
### Extensions (as documentation) must be added as `protocol_<name>.md`.

---

## Bool(ean)s
##### Type: `bool`.

### Bool is a boolean where `0` represents false and `255` represents true.
```go
[0]     // is false
[255]   // is true
```

## Numbers
##### Types: `int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64`.

### Numbers as described are variable uintegers (VarUint) implemented by Go's standard library.

```go
[64]                                        // 64
[128 8]                                     // 1024
[255 255 255 255 255 255 255 255 255 1]    // 18446744073709551615
[129 128 128 128 128 128 128 128 128 1]     // -9223372036854775807
```

## Floats and Complexes
##### Types: `float32, float64, complex64, complex128`.

### Very like numbers floats are a representation made by `math.Float32bits` or `math.Float64bits` then parsed as an uint32 or uint64.
### Complex numbers are two floats, the first being the real part, and the second part being the imaginary part.

```go
[159 129 137 141 171 243 160 138 69] // 6.2e+24
[128 128 176 183 177 178 195 128 67 145 212 183 137 152 243 229 192 61] // (6e+14+2e-12i)
```

## Arrays
##### Types: `[size]T`.

### Arrays considers the size of the variable, during encoding it might place zeroes to make sure it is valid.
### An array value is parsed as its underlying type.

```go
[13 64 128 8] // [3]uint64{13, 64, 1024}
```

## Slices
##### Types: `[]T`.

### As the same as array but it first includes the size of slice.

```go
[3 13 64 128 8] // []uint64{13, 64, 1024}
```

## Strings
##### Types: `string`.

### Strings are parsed just as slices, size first then data.

```go
[13 72 101 108 108 111 44 32 87 111 114 108 100 33] // Hello, World!
```

## Maps
##### Types: `map[K]V`.

### Maps are parsed as first size then its keys and values as how their respective types are handled.

```go
[2 5 72 101 108 108 111 6 87 111 114 108 100 33 3 66 105 110 8 65 119 101 115 111 109 101 33] // map[Bin:Awesome! Hello:World!]
```

## Struct
##### Types: `T (as struct)`

### A tag is an identifier for a field in a structure. The same packet can be used in different structures as long as they match the field tag.
### Structs are parsed as tag first then data handled as their type.


## Tag
#### Defaults to field number in structure starting from 1.
- Go: Following a struct field place `bin:"<number>""`.
- JS: Inside a class add a static field `BIN_TAG` as an object and place the field name and the id `static BIN_TAG = {<field>: <number>}`.

```go
// struct { Hello string `bin:"10"`; Bin string `bin:"20"` }
[10 6 87 111 114 108 100 33 20 8 65 119 101 115 111 109 101 33] // {World! Awesome!}
```
