package main

import "fmt"

func main() {
	var token string

	tokenValid := "12345abcde"

	fmt.Print("Masukan Token: ")
	fmt.Scan(&token)

	for token != tokenValid {
		fmt.Print("Masukan Token: ")
		fmt.Scan(&token)
	}
	fmt.Println("Selamat anda berhasil login")
}