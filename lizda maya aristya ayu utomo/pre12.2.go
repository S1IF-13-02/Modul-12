package main

import "fmt"

func main() {
	const tokenBenar = "12345abcde"
	var input string

	for {
		fmt.Print("masukkan token: ")
		fmt.Scan(&input)

		if input == tokenBenar {
			fmt.Println("Selamat Anda berhasil login")
			break
		}
	}
}
