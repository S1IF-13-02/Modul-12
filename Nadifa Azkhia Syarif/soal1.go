package main

import "fmt"

func main() {
	const userBenar = "Admin"
	const passBenar = "Admin"

	var u, p string
	gagal := 0

	fmt.Scan(&u, &p)

	for u != userBenar || p != passBenar {
		gagal++
		fmt.Scan(&u, &p)
	}

	fmt.Println(gagal, "percobaan gagal login")
}
