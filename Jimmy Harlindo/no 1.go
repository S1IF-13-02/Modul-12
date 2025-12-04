package main

import (
	"fmt"
)

func main() {
	const userBenar = "Admin"
	const passBenar = "Admin"

	var user, pass string
	gagal := 0

	// while loop versi Go → for tanpa syarat
	for {
		fmt.Print("Masukkan username: ")
		fmt.Scan(&user)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&pass)

		// cek benar atau tidak
		if user == userBenar && pass == passBenar {
			break // keluar dari loop jika berhasil login
		}

		gagal++ // hitung gagal login
		fmt.Println("Username atau password salah, coba lagi.")
	}

	fmt.Printf("%d percobaan gagal login\n", gagal)
}
