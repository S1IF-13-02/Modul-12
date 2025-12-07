package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan dua bilangan x dan y: ")
	fmt.Scan(&x, &y)

	if y == 0 {
		fmt.Println("Tidak bisa dibagi dengan nol")
		return
	}

	hasil := 0

	for x >= y {
		x = x - y
		hasil++
	}

	fmt.Println(hasil)
}
