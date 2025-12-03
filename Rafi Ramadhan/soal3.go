ackage main

import "fmt"

func main() {
    var x, y int
    fmt.Print("Masukkan dua bilangan (x y): ")
    fmt.Scan(&x, &y)

    h := 0

    for x >= y {
        x = x - y   
        h++   
    }
