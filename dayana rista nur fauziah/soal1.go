package main

import "fmt"

func main() {
	var user, pass string
	gagal := 0

	for user != "Admin" || pass != "Admin" {
		fmt.Scan(&user, &pass)

		if user != "Admin" || pass != "Admin" {
			gagal++ // sama aja kaya (gagal=gagal+1)
		}
	}

	fmt.Println(gagal, "percobaan gagal login")
}
