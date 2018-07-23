
// A defer statment defers the execution of a function until the surrounding function 
// returns.

// The deferred call's arguments are evaluated immediately, but the funciton call is no
// executed until the surrounding function returns


package main

import "fmt"


func main() {
    defer fmt.Println("World")
    fmt.Println("hello")
}
