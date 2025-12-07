package main

import "fmt"

func main() {
	var token string
	var tokenValid = "12345abcde"

	fmt.Print("Masukkan Token : ")
	fmt.Scan(&token)

	for token != tokenValid {
		fmt.Print("Masukkan Token : ")
		fmt.Scan(&token)
	}
	fmt.Println("Selamat Anda Berhasil Login")
}
