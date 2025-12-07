package main

import "fmt"

func main() {
	var bilangan, digit int

	fmt.Print("Masukkan bilangan = ")
	fmt.Scan(&bilangan)

	for bilangan > 0 {
		digit = bilangan % 10
		fmt.Println(digit)
		bilangan = bilangan / 10
	}
}