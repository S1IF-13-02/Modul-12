package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)

	for n > 0 {
		digit := n % 10
		fmt.Println(digit)
		n = n / 10
	}
}
