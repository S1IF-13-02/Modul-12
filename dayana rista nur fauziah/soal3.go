package main

import "fmt"

func main() {
	var y, sisa int

	fmt.Print("Masukkan nilai x dan y: ")
	fmt.Scan(&sisa, &y)

	hasil := 0

	for sisa >= y {
		sisa = sisa - y
		hasil = hasil + 1
	}

	fmt.Println("Hasil : ", hasil)
}
