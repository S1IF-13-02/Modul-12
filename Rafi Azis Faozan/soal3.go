package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)
	hasil := 0
	for x >= y {
		x = x - y
		hasil = hasil + 1
	}
	fmt.Println("Hasil x div y: ", hasil)
}
