package main

import "fmt"

func main() {
	var n int
	fmt.Print("Massukkan bilangan bulat: ")
	fmt.Scan(&n)
	a := 0
	b := 1
	j := 0
	for j < n {
		fmt.Print(a, " ")
		temp := a + b
		a = b
		b = temp
		j = j + 1
	}
}
