package main

import "fmt"

func main() {
	var n int
	var a, b int = 0, 1
	var count int = 0
	var next int

	fmt.Print("Masukkan jumlah bilangan Fibonacci: ")
	fmt.Scan(&n)

	for count < n {
		if count == 0 {
			fmt.Print(a)
		} else if count == 1 {
			fmt.Print(" ", b)
		} else {
			next = a + b
			fmt.Print(" ", next)
			a = b
			b = next
		}
		count++
	}
	fmt.Println()
}
