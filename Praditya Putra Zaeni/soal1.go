package main

import "fmt"

func main() {
    const benarUser = "Admin"
    const benarPass = "Admin"

    var user, pass string
    percobaanGagal := 0

    for {
        fmt.Scan(&user, &pass)

        if user == benarUser && pass == benarPass {
            break
        } else {
            percobaanGagal++
        }
    }

    fmt.Printf("%d percobaan gagal login\n", percobaanGagal)
}
