package main

import "fmt"

func main() {
	var n, i int
	fmt.Scan(&n)

	i = n

	for i > 1 {
		fmt.Print(i, " x ")
		i = i - 1
	}
	fmt.Print(1)
}
