package main

import "fmt"

func main() {
    valid := "12345abcde"
    var input string

    fmt.Print("Masukkan token: ")
    fmt.Scanln(&input)

    for input != valid {  
        fmt.Print("Masukkan token: ")
        fmt.Scanln(&input)
    }
    fmt.Println("Selamat Anda berhasil login")
}
