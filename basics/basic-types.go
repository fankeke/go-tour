// Go's basic types are
// bool string int int8 int16 int32 int64 
//             uint uint8 uint16 uint 32 uint64 uintptr
// byte (alias for uint8)
// rune (alias for int32)

// The int, uint and uintptr types are usually 32 bits wide on 32-bit system and 64 bits wide on 64-bit systems. when you need an interger value you should use int unless you have a specific reason to use a sized or unsigned integer type.


package main

import (
    "fmt"
    "math/cmplx"
    )

var (
        ToBe    bool = false
        MaxInt  uint64 = 1<<64 - 1
        z complex128 = cmplx.Sqrt(-5 + 12i)
    )

func main() {
    fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)
}

