package main

import "fmt"

func main() {
	userValid := "Admin"
	passValid := "Admin"
	var username, password string
	n := 0
	
	fmt.Print("Masukan Username dan Password : ")
	fmt.Scan(&username, &password)
	for username != userValid || password != passValid {
		n++
		fmt.Println(n, "percobaan gagal login")
		fmt.Print("Masukan Username dan Password : ")
		fmt.Scan(&username, &password)
	}
	fmt.Println("Total percobaan gagal :", n)
	fmt.Println("Anda Berhasil Login")
}
