# Bin Protocol Extension: Interface
### This file describes the behaviour of `interface{}` type with the Bin Protocol.

---

## The `bin.Interface` function.
### This function is responsible to turn any type into `interface{}`.
- Basic types -> `interface{}`.
- `[size]T` -> `[size]interface{}`
- `[]T` -> `[]interface{}`.
- `map[K]V` -> `map[interface{}]interface{}`
- `T (as struct)` -> `interface{}` as a representation of fields as `interface{}`.

## *Struct
### This is a map (`map[int]interface{}`) representation of a struct, it can be visualized by calling `Map()` or parse to a struct with `As()`.

---

## Basic Types
##### Types: `bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128, string`.

### When encoding those types, we append their respective kind as in `reflect.Kind`.

```go
[1 255] // true
[2 128 32] // 4096
[14 225 245 209 240 250 168 216 149 64] // 13.69
[16 128 128 128 128 128 128 128 128 64 128 128 128 128 128 128 128 136 64] // (2+4i)
[24 13 72 101 108 108 111 44 32 87 111 114 108 100 33] // Hello, World!
```

## Arrays and Slices
##### Type: `[size]T, []T`

### Encoded as normal, type is set to `interface{}` and parsed as its underlying type.

```go
[2 16 64] // [16 64] (as []uint64)
[23 1 0 11 2 16 64] // [16 64] (as interface{}, underlying []uint64)
[23 1 0 20 2 11 16 11 64] // [16 64] (as interface{}, underlying []interface{})
```

## Maps
##### Type: `map[K]V`

### The map type is set to `interface{}`, the key and value types may be set to `interface{}` too but will remain its underlying type

```go
[21 24 2 1 3 66 105 110 10] // map[Bin:10] (map[string]int)
[21 20 20 1 24 3 66 105 110 2 10] // map[Bin:10] (map[interface{}]interface{})
```

## Struct
##### Types: `T (as struct)`

### A struct will be encoded as length of tags then tags followed by the type and data.
### The requested struct will try to set fields according to the tags available.

```go
// struct { Hello string `bin:"10"`; Bin string `bin:"20"` }
[25 2 10 24 6 84 104 101 114 101 33 20 24 4 78 105 99 101] // {There! Nice}
```