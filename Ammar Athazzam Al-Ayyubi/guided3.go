package main

import "fmt"

func main() {
	var a, b, c, i, n int
	fmt.Print("Masukkan Bilangan: ")
	fmt.Scan(&n)
	
	a = 0
	b = 1
	fmt.Print(a, " ", b, " ")

	for i = 2; i < n; i++ {
		c = a + b
		fmt.Print(c, " ")
		a = b
		b = c
	}
	fmt.Println()
}
