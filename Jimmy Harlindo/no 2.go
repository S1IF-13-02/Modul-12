package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)

	// while loop di Go menggunakan: for n > 0 { ... }
	for n > 0 {
		digit := n % 10     // ambil digit terakhir
		fmt.Println(digit)  // tampilkan digit
		n = n / 10          // hapus digit terakhir
	}
}
