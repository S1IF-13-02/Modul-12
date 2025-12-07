package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const correctUsername = "Admin"
	const correctPassword = "Admin"
	var failedAttempts int

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Masukkan username: ")
		usernameInput, _ := reader.ReadString('\n')
		usernameInput = strings.TrimSpace(usernameInput)

		fmt.Print("Masukkan password: ")
		passwordInput, _ := reader.ReadString('\n')
		passwordInput = strings.TrimSpace(passwordInput)

		if usernameInput == correctUsername && passwordInput == correctPassword {
			break
		} else {
			fmt.Println("Login gagal. Silakan coba lagi.\n")
			failedAttempts++
		}
	}

	fmt.Printf("Login berhasil setelah %d kali percobaan gagal.\n", failedAttempts)
}
