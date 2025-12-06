package main

import "fmt"

func main() {
	var n string
	for {
		fmt.Print("masukan token : ")
		fmt.Scan(&n)

		if n == "12345abcde" {
			fmt.Print("selamat anda berhasil login")
			break
		} else {
			fmt.Print("token yang anda masukan salah")
		}
	}
}
