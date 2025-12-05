package main

import "fmt"

func main() {
	var usn string
	var password string
	fmt.Print("Masukkan username: ")
	fmt.Scan(&usn)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)
	j := 0
	if usn == "Admin" {
		for password != "Admin" {
			fmt.Print("Masukkan password: ")
			fmt.Scan(&password)
			j = j + 1
		}
	}
	fmt.Print(j, " percobaan gagal login")
}
