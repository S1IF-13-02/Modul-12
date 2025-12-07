package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	for n > 0 {
		var digit int
		digit = n % 10     
		fmt.Println(digit) 
		n = n / 10         
	}
}