package main

import "fmt"

func main() {
	var user, pass string
	gagal := 0

	berhasilLogin := false

	for !berhasilLogin {
		fmt.Print("Masukkan Username dan Password: ")
		fmt.Scan(&user, &pass)

		if user == "Admin" && pass == "Admin" {
			berhasilLogin = true
		} else {
			gagal++
		}
	}

	fmt.Printf("%d percobaan gagal login\n", gagal)
}
