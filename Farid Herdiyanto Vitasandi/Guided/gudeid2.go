package main

import "fmt"

func main() {

var password string
const token = "12345abcde"

for password != token {
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	if password != token {
		fmt.Println("Password salah, silahkan coba lagi.")
	}
}

fmt.Println("Selamat Anda berhasil login")
}
