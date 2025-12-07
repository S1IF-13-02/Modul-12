package main

import "fmt"


func main() {
    var n int
    fmt.Print("Masukkan angka: ")
    fmt.Scan(&n)

    if n == 0 || n == 1 {
        fmt.Println("1")
        return
    }

    i := n
    for { 
        if i == 1 {
            fmt.Print("1")
            break
        }
        fmt.Printf("%d x ", i)
		i--
}
}