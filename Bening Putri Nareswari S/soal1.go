package main

import (
	"fmt"
)

func main() {
	const (
		validUser = "Admin"
		validPass = "Admin"
	)

	var hitunganGagal int = 0
	var loginSukses bool = false

	for loginSukses == false {
		var userIn string
		var passIn string

		fmt.Print("Input Username: ")
		fmt.Scanln(&userIn)

		fmt.Print("Input Password: ")
		fmt.Scanln(&passIn)

		if userIn == validUser && passIn == validPass {
			loginSukses = true
			fmt.Println("Login Berhasil")
		} else {
			hitunganGagal++
			fmt.Printf("Gagal. Coba lagi.\n")
		}
	}

	fmt.Printf("Total: %d percobaan gagal login", hitunganGagal)
}
