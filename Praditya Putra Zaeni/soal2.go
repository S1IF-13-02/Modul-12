package main

import "fmt"

func main() {
    var bilangan int
    fmt.Scan(&bilangan)

    for bilangan > 0 {
        digit := bilangan % 10
        fmt.Println(digit)
        bilangan = bilangan / 10
    }
}
