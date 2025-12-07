package main

import "fmt"

func main() {
    tokenValid := "12345abcde"
    var token string
	
    fmt.Print("Masukkan token: ")
    fmt.Scan(&token)

    for token != tokenValid {
        fmt.Println("Token salah, coba lagi.\n")
        fmt.Print("Masukkan token: ")
        fmt.Scan(&token)
    }
	fmt.Println("Login berhasil!")
}
