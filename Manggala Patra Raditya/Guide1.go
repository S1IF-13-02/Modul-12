package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukan: ")
    fmt.Scan(&n)

    if n == 0 {
        fmt.Println(1)
        return
    }

    for i := n; i >= 1; i-- {
        if i == 1 {
            fmt.Print(i)
        } else {
            fmt.Print(i, " x ")
        }
    }

    fmt.Println()
}