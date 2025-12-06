package main

import "fmt"

func main() {
	var user, pass string
	gagal := 0

	for {
		fmt.Scan(&user, &pass)

		if user == "Admin" && pass == "Admin" {
			break
		}

		gagal++
	}

	fmt.Printf("%d percobaan gagal login\n", gagal)
}
