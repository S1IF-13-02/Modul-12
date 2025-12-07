package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan x (x >= y): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan y: ")
	fmt.Scan(&y)

	if x < y || x <= 0 || y <= 0 {
		fmt.Println("Input tidak valid. Pastikan x >= y dan keduanya positif.")
		return
	}
	hasil := 0
	for x >= y {
		x = x - y
		hasil++
	}

	fmt.Println("Hasil integer division (x div y):", hasil)
}
