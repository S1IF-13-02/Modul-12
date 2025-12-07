package main

import "fmt"

func main() {
	var username, pass string
	var gagal int

	for {
		fmt.Scan(&username, &pass)

		if username == "Admin" && pass == "Admin" {
			break
		}

		gagal++
	}

	fmt.Println(gagal, "percobaan gagal login")
}
