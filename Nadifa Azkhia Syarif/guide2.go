package main

import "fmt"

func main() {
	const tokenBenar = "12345abcde"
	var input string

	fmt.Print("Masukkan token: ")
	fmt.Scan(&input)

	for input != tokenBenar {
		fmt.Print("Masukkan token: ")
		fmt.Scan(&input)

	}
	fmt.Println("Selamat anda berhasil login")
}
