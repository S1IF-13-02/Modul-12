package main

import "fmt"

func main() {
	const ValidToken = "12345abcde" 
	var token string

	for { 
		fmt.Print("Masukkan token: ")
		fmt.Scan(&token)

	if token == ValidToken {
		fmt.Print("Selamat Anda Berhasil Login")
		break
		}
	}
}

