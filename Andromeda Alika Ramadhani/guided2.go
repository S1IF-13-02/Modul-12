package main

import "fmt"

func main() {
	var token string

	for token != "12345abcde" {
		fmt.Print("Masukkan Password: ")
		fmt.Scan(&token)
	}
	fmt.Println("Selamat Anda berhasil login")
}
