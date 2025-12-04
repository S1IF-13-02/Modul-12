package main

import "fmt"
func main() {
	var password string
	passwordValid := "12345abcde"

	fmt.Scan(&password)

	for password != passwordValid {
		fmt.Scan(&password)
	}  
	fmt.Println("Selamat Anda Berhasil Masuk" )
}
	