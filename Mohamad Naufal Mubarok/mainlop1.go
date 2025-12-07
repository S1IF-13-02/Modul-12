package main

import "fmt"

func main() {
	var user, pass string
	correctUser := "Admin"
	correctPass := "Admin"
	gagal := 0

	fmt.Print("Masukkan username: ")
	fmt.Scan(&user)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&pass)

	for user != correctUser || pass != correctPass {

		gagal++
		fmt.Println("Username atau password salah! Coba lagi.\n")

		fmt.Print("Masukkan username: ")
		fmt.Scan(&user)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&pass)
	}

	fmt.Println(gagal, "percobaan gagal login")
}