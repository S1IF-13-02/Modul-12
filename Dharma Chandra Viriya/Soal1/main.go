package main

import "fmt"

func main() {
	var usr, pwd string

	counter := 0
	for usr != "Admin" && pwd != "Admin" {
		fmt.Print("Masukkan username: ")
		fmt.Scanln(&usr)
		fmt.Print("Masukkan password: ")
		fmt.Scanln(&pwd)
		fmt.Println("Username atau Password Yang Anda Masukkan Salah")
		counter++
	}

	fmt.Printf("%d percobaan gagal login\n", counter-1)
}
