package main

import "fmt"

func main() {
	var number int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&number)

	for number > 0 { 
		digit := number % 10
		fmt.Println(digit)
		number /= 10
	}
	fmt.Println()
}