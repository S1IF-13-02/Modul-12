package main

import "fmt"

func main() {
	const usernamebenar, passwordbenar = "Admin", "Admin"

	var username, password string
	gagallogin := 0

	for {
		fmt.Print("masukan username : ")
		fmt.Scan(&username)
		fmt.Print("masukan password : ")
		fmt.Scan(&password)

		if username == usernamebenar && password == passwordbenar {
			break
		}

		gagallogin = gagallogin + 1
	}

	fmt.Printf("%d percobaan gagal login\n", gagallogin)
}
