package main

import "fmt"

func main() {
	var token string
	validToken := "12345abcde"

	fmt.Print("Masukkan token: ")
	fmt.Scan(&token)

	for token != validToken {
		fmt.Println("Token salah! Silakan coba lagi.")
		fmt.Print("Masukkan token: ")
		fmt.Scan(&token)
	}

	fmt.Println("Selamat anda berhasil login")
}
