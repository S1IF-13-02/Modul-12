package main

import "fmt"

func main() {
	
	var bilangan int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&bilangan)

	for bilangan < 0 {
		fmt.Println("Bilangan harus positif!")
		fmt.Print("Masukkan bilangan bulat positif: ")
		fmt.Scan(&bilangan)
	}

	hitungDigit := 0

	for bilangan > 0 {
		digit := bilangan % 10
		hitungDigit++ 

		fmt.Printf("Digit ke-%d: %d\n", hitungDigit, digit)

		bilangan = bilangan / 10
	}
}
