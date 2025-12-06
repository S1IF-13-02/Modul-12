package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	if N >= 2 {
		a, b := 0, 1
		fmt.Print(a, " ", b)

		for i := 2; i < N; i++ {
			next := a + b
			fmt.Print(" ", next)
			a, b = b, next
		}

		fmt.Println()
	}
}
