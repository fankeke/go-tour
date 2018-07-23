/*
   There are two reasones to use a pointer receiver.

   The first is so that the method can modify the value that its receiver points to


   The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct.

   In this example, both Scal and Abs are with receiver type *Vertex, even though the Abs method need't modify its receiver.

   in general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

*/

package main

import (
     "fmt"
     "math"
     )

type Vertex struct {
    X, Y float64 
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
    v := &Vertex{3,4}
    fmt.Println(v.Abs())
    
    v.Scale(5)

    fmt.Println(v.Abs())
}


