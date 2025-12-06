package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	a := 0
	b := 1

	fmt.Print(a, " ", b)

	i := 2

	for i < n {
		c := a + b
		fmt.Print(" ", c)

		a = b
		b = c
		i++
	}
}
