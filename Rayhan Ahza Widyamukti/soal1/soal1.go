package main

import "fmt"

func main() {
	var username string
	var password string
	var gagal int
	var loginBenar bool

	for loginBenar == false {

		fmt.Println("Masukkan username dan password: ")
		fmt.Scan(&username, &password)

		if username == "Admin" && password == "Admin" {
			loginBenar = true
		} else {
			gagal = gagal + 1
			fmt.Println("Username atau password salah!")
		}
	}

	fmt.Printf("%d percobaan gagal login\n", gagal)
}
