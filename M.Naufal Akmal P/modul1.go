package main

import "fmt"

func main() {
    var i int

	fmt.Scan(&i)
	
    hasil := 1

    for i >= 1 {
        fmt.Print(i)
        if i != 1 {
            fmt.Print(" x ")
        }
        hasil *= i
        i--
    }

    fmt.Println(" =", hasil)
}