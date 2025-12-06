package main

import "fmt"

func main() {
    var user string
    var password string

    userbenar := "Admin"
    passbenar := "Admin"

    gagal := 0

    fmt.Scan(&user, &password)

    for user != userbenar || password != passbenar {
        gagal++

        fmt.Scan(&user, &password)
    }

    fmt.Println(gagal, "percobaan gagal login")
}