package main

import "fmt"

func main() {
	var u, p string
	var gagal int
	const token = "Admin"

	for {
		fmt.Print("Masukkan User dan Password: ")
		fmt.Scan(&u, &p)

		if u != token || p != token {
			gagal++
			continue
		}
		if u == token && p == token {
			fmt.Println(gagal, "percobaan gagal login")
			break
		}
	}
}
