package main

import "fmt"

func main() {
    var n int

    fmt.Print("masukan nilai: ")
    fmt.Scan(&n)

    if n < 2 {
        fmt.Println("masukan n minimal 2")
        return
    }

    f1 := 0
    f2 := 1

    fmt.Print(f1, " ", f2)
    for i := 3; i <= n; i++ {
        fn := f1 + f2
        fmt.Print(" ", fn)
        f1 = f2
        f2 = fn
    }

    fmt.Println()
}
