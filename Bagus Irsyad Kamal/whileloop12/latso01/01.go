package main

import (
	"fmt"
)

func main() {
	benarUser := "Admin"
	benarPass := "Admin"

	var user, pass string
	gagal := 0
	fmt.Scan(&user, &pass)

	for user != benarUser || pass != benarPass {
		gagal++
		fmt.Scan(&user, &pass)
	}

	fmt.Println(gagal, "percobaan gagal login")
}
