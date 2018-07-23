package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}
var m = map[string]Vertex{
    "Bell Labs" : Vertex{
        40, 10,
    },
    "Google": Vertex{
        30, 20,
    },
}
func main() {
    fmt.Println(m)
}
