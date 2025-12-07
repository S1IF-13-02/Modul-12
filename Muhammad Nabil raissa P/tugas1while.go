package main

import "fmt"

func main() {
    const validUser, validPass = "Admin", "Admin"

    var username, password string
    gagal := 0

    for {
        fmt.Scan(&username, &password)

        if username == validUser && password == validPass {
            break 
        }
        
        gagal++ 
    }

    fmt.Printf("%d percobaan gagal login\n", gagal)
}