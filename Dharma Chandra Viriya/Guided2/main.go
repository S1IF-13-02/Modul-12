package main

import "fmt"

func main() {
	const validToken string = "12345abcde"
	var token string

	fmt.Print("Masukan token: ")
	fmt.Scan(&token)

	for token != validToken {
		fmt.Println("Gagal Login Token Anda Salah")

		fmt.Print("Masukan token: ")
		fmt.Scan(&token)
	}

	fmt.Println("Selamat Anda Berhasil Login")

}
