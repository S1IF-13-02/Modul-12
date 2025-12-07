package main

import "fmt"

func main() {
	var user, pass string
	var salah int

	fmt.Scan(&user, &pass)

	for user != "Admin" || pass != "Admin" {
		salah = salah + 1
		fmt.Scan(&user, &pass)
	}

	fmt.Println(salah, "percobaan gagal login")
}
