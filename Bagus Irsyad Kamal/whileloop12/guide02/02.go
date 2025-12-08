package main  
import "fmt"  
func main() {  
    var token string  
	fmt.Print("Masukkan Kode Password : ")
    fmt.Scan(&token) 
    for token != "12345abcde" { 
        fmt.Scan(&token) 
    } 
    fmt.Println("Selamat Anda berhasil login") 
} 