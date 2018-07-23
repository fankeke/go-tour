
// defered function calls are pushed onto a stack. when a function returns
//its defered calls are executed in last-in-first-out order.

package main

import "fmt"

func main() {
    fmt.Println("counting")


    for i := 0; i < 10; i++ {
        defer fmt.Println(i) //push function calls into stack
    }

    fmt.Println("done")
}

