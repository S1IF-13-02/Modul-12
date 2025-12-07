package main

import "fmt"
func main() {
    var n int
    fmt.Print("Masukan angka bilangan bulat positif : ")
    fmt.Scan(&n)

    if n < 0 {
        fmt.Println("Tidak bisa bilangan negatif")
        return
    }

    total := 1
    for n > 0 {
        total *= n
        fmt.Print(n)
        if n > 1 {
            fmt.Print(" x ")
        }
		n--
    }

    fmt.Println(" =", total)
}
