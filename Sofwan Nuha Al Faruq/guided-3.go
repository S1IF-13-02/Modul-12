package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan jumlah bilangan Fibonacci: ")
	fmt.Scan(&n)

	a, b := 0, 1
	count := 0

	for count < n {
		if count == 0 {
			fmt.Print(a)
		} else if count == 1 {
			fmt.Print(" ", b)
		} else {
			next := a + b
			fmt.Print(" ", next)
			a = b
			b = next
		}
		count++
	}
	fmt.Println()
}
