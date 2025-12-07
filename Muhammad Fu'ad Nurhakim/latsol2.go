package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&angka)

	if angka < 0 {
		angka = -angka
	}

	if angka == 0 {
		fmt.Println(0)
		return
	}

	for angka > 0 {
		digit := angka % 10
		fmt.Println(digit)
		angka = angka / 10
	}
}
