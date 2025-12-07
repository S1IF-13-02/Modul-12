package main

import "fmt"

func main() {
	var token string
	validToken := "12345abcde"

	for {
		fmt.Print("Masukkan token: ")
		fmt.Scanln(&token)

		if token == validToken {
			fmt.Println("Selamat Anda berhasil login")
			break
		}
	}
}
