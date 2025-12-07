package main

import "fmt"

func main() {
	var x, y int
	var hasil int

	fmt.Scan(&x, &y)

	for x >= y {
		x = x - y
		hasil = hasil + 1
	}

	fmt.Println(hasil)
}
