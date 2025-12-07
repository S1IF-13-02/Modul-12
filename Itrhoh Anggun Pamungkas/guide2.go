package main

import "fmt"

func main() {
	var input string

	tokenBenar := "12345abcde"
	for input != tokenBenar {
		fmt.Print("token: ")
		fmt.Scan(&input)
	}
	fmt.Println("Selamat Anda berhasil login")
}
