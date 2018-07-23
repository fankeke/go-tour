package main

import (
    "fmt"
    "math"
)

// Methods with pointer receivers can modify the value to which the receive 
//pointer

// Since methods often need to modify their receiver, pointer receivers 
//are more common than value receivers


// Functions with a pointer argument must take a pointer:
/* 
   var v Vertex
   ScaleFunc(v, 5) // Compile error
   ScaleFunc(&v, 5) // ok

   while methods with pointer receivers take either a value or a pointer all works well

   var v Vertex
   v.Scale(5) //ok
   
   p := &v
   p.Scale(10) //ok
*/

type Vertex struct {
    X, Y float64
}

func(v *Vertex)Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3,4}
    v.Scale(10)
    fmt.Println(v.Abs())
}

