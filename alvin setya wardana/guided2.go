package main
import "fmt"
func main() {
    var input string
    const password = "1234abcde"
    for {
        fmt.Print("masukkan password: ")
        fmt.Scanln(&input)
        if input == password {
            fmt.Println("login berhasil")
            break
        } else {
            fmt.Println("password salah, coba lagi")
        }
    }
}
