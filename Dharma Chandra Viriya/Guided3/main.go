package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukan n: ")
	fmt.Scan(&n)

	i := 2
	a := 0
	b := 1

	fmt.Print(a, " ")
	fmt.Print(b, " ")

	for i < n {
		temp := a + b
		fmt.Print(temp, " ")

		a = b
		b = temp

		i++
	}
	fmt.Println()
}
