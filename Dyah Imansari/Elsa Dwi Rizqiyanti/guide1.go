package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	if n == 0 || n == 1 {
		fmt.Println("1")
		return
	}

	i := n
	for i >= 1 {
		if i == 1 {
			fmt.Print("1")
		} else {
			fmt.Print(i, " x ")
		}
		i--
	}
}
