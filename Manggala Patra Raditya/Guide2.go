package main

import "fmt"

func main() {
    validToken := "12345abcde"
    var input string

    for {
        fmt.Print("Masukan token: ")
        fmt.Scan(&input)

        if input == validToken {
            fmt.Println("Selamat Anda berhasil login")
            break
        }
    }
}