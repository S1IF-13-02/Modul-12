package main

import "fmt"

func main() {
	var bil int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&bil)

	for bil > 0 {
		digit := bil % 10
		fmt.Println(digit)
		bil = bil / 10
	}
}
