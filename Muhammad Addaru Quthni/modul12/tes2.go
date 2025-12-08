package main
import "fmt"
func main() {
	var a string = "12345abcde"
	fmt.Print("Masukkan token: ")
	fmt.Scan(&a)
	for a != "12345abcde" {
		fmt.Print("")
		fmt.Scan(&a)
	}
	fmt.Println("Selamat anda berhasil masuk")
}