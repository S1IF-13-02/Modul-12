package main

import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)

    hasil := 0

    for x >= y {
        x = x - y   
        hasil = hasil + 1   
    }
	fmt.Println(hasil)
}