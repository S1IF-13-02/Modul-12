package main

import "fmt"

func main() {
	var digit int

	fmt.Scan(&digit)
	
	for digit > 0 {
		sisa := digit % 10
		fmt.Println(sisa)
		digit = digit / 10
	}
}