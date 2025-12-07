package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan Bilangan Bulat: ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("N harus lebih atau sama dengan 2")
		return
	}
	a := 0
	b := 1
	count := 0

	for count < n {
		fmt.Print(a, " ")
		a, b = b, a+b
		count++
	}
}
