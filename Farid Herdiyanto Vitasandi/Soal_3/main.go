package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan y: ")
	fmt.Scan(&y)

	for y == 0 || x < y {
        if y == 0 {
            fmt.Println("y tidak boleh 0!")
			fmt.Print("Masukkan y: ")
            fmt.Scan(&y)
        }
        if x < y {
            fmt.Println("x harus lebih besar dari y!")
			fmt.Print("Masukkan x: ")
            fmt.Scan(&x)
        }
    }
    
    hasil := 0
    sisa := x
    
    for sisa >= y {
        sisa -= y
        hasil++
    }
    
    fmt.Printf("Hasil: %d div %d = %d\n", x, y, hasil)
}
