package main

import "fmt"

func main() {
	var token string

	for token != "12345abcde"{
		fmt.Print("Masukkan password = ")
		fmt.Scan(&token)
	}
	fmt.Print("Selamat anda berhasil login")
}