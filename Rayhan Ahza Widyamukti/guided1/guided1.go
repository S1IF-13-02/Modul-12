package main

import "fmt"

func main() {
	var n, i int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println(1)
	} else {
		for i = n; i > 0; i-- {
			fmt.Print(i)
			if i > 1 {
				fmt.Print(" * ")
			}
		}
		fmt.Println()
	}
}
