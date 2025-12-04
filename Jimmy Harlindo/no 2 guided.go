package main

import "fmt"

func main(){
	const tokenBenar = "12345abcde"
	var input string

	for {
		fmt.Print("Masukkan token: ")
		fmt.Scanln(&input)

		if input == tokenBenar {
			fmt.Println("Selamat Anda berhasil login")
			break
		}
	}
}
