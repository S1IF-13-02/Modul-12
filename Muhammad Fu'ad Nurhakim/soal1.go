package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)
	for a == 0 {
		fmt.Println(1)
		return
	}
	fmt.Print(a)
	for a > 1 {
		a--
		fmt.Print(" x ", a)
	}
}
