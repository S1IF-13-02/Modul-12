package main

import "fmt"

func main() {
	var n, i, a, b, hasil int
	fmt.Scan(&n)

	a = 0
	b = 1
	i = 0

	for i < n {
		fmt.Print(a, " ")
		hasil = a + b
		a = b
		b = hasil
		i++
	}
}
