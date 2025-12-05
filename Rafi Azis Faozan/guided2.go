package main

import "fmt"

func main() {
	var token string
	fmt.Print("Masukkan token masuk: ")
	fmt.Scan(&token)
	for token != "12345abcde" {
		fmt.Print("Masukkan token masuk: ")
		fmt.Scan(&token)
	}
	fmt.Print("Selamat anda berhasil login")
}
