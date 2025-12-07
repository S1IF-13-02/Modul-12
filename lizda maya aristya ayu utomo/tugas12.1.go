package main

import "fmt"

func main() {
	const userbenar = "Admin"
	const passbenar = "Admin"

	var username, password string
	gagal := 0

	fmt.Scan(&username, &password)

	for username != userbenar || password != passbenar {
		gagal++
		fmt.Scan(&username, &password)
	}
	fmt.Print(gagal, "percobaan gagal login")
}
