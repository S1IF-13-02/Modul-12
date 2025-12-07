package main

import "fmt"

func main() {
    var N int

    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&N)
    a, b := 0, 1
    fmt.Print("bilangan fibonacci:")
    for i := 0; i < N; i++ {
        fmt.Printf("%d ", a)
        a, b = b, a+b

    }
    fmt.Println()
}
