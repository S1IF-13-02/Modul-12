package main

import (
	"fmt"
)

func main() {
	var x int
	var y int
	var hasilBagi int = 0

	fmt.Print("Masukkan bilangan X (pembilang): ")
	fmt.Scanln(&x)

	fmt.Print("Masukkan bilangan Y (pembagi): ")
	fmt.Scanln(&y)

	for x >= y {
		x = x - y
		hasilBagi++
	}

	fmt.Printf("Hasil integer division (X div Y) adalah: %d\n", hasilBagi)
}
