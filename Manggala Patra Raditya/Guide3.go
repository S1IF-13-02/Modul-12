package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukan: ")
    fmt.Scan(&n)

    a := 0
    b := 1

    fmt.Print(a, " ", b)

    for i := 3; i <= n; i++ {
        c := a + b
        fmt.Print(" ", c)
        a = b
        b = c
    }

    fmt.Println()
}