package main

import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)

    hasil := 0
    nilai := x

    for nilai >= y {
        nilai = nilai - y
        hasil++
    }

    fmt.Println(hasil)
}