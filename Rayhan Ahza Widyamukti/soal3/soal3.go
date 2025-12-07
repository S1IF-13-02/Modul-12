package main

import "fmt"

func main() {
    var x int
    var y int
    var hasil int

    fmt.Print("Masukkan x: ")
    fmt.Scan(&x)

    fmt.Print("Masukkan y: ")
    fmt.Scan(&y)

    for x >= y {
        x = x - y   
        hasil = hasil + 1
    }

    fmt.Println(hasil)
}
