package main

import "fmt"

//fibnacci is a function that returns
// a function that returns an int.

func fibonacci() func() int {
    i, j := 1 ,0
    return func() int {
        i, j = j, i+j
        return i
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
