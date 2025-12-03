package main
import "fmt"
func main() {
    var n, j int

    fmt.Print("masukan biangan : ")
    fmt.Scan(&n)

    j = n
    for j > 1 {
        fmt.Print(j, " x ")
        j--
    }
    fmt.Println(1)
}
