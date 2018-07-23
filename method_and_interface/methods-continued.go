package main

// You can only declare a method with a receiver whose type is defined 
// in the same package as the method.
// You cannot declare a method with a receiver whose type is define in another
// package ( which inclus the built-in types such as int)

import (
        "fmt"
        "math"
)

type MyFloat float64

func(f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    } 

    return float64(f)
}

func main() {
    f := MyFloat(-math.Sqrt(2))
    fmt.Println(f.Abs())
}
