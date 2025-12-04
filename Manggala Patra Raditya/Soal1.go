package main

import "fmt"

func main() {
    usernameBenar := "Admin"
    passwordBenar := "Admin"

    var u, p string
    gagal := 0

    for {
        fmt.Print("Masukan username dan password: ")
        fmt.Scan(&u, &p)

        if u == usernameBenar && p == passwordBenar {
            fmt.Println(gagal, "percobaan gagal login")
            break
        } else {
            gagal++
        }
    }
}