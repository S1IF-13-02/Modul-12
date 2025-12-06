package main

import "fmt"

func main() {
	var t string

	for {
		fmt.Print("masukan pass: ")
		fmt.Scan(&t)
		if t == "12345abcde" {
			fmt.Print("selamat anda berhasil login")
			break
		} else {
			fmt.Println("coba lagi")
		}
	}
}
