package main

import "fmt"

func main() {
	var username, password string
	failedAttempts := 0
	correctUsername := "Admin"
	correctPassword := "Admin"

	for username != correctUsername || password != correctPassword {
		fmt.Print("Masukkan username: ")
		fmt.Scanln(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scanln(&password)

		if username != correctUsername || password != correctPassword {
			failedAttempts++
		}
	}

	fmt.Printf("%d percobaan gagal login\n", failedAttempts)
}