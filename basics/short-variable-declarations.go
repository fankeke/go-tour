// inside a function, the `:=` short assignment statement can be used in
// place of a `var` declaration with implicit type.

// outside a function, every statement begins with a keywork (`var`, `func` and so on) and so the `:=` constuct is no availabe


package main

import "fmt"

func main() {
    var i, j int = 1, 2
    k := 3

    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)
}
