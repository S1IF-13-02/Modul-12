package main
import "fmt"

func main() {
	var n string
	const token = "12345abcde"
	

	for {
		fmt.Print("Masukkan Password: ")
		fmt.Scan(&n)
		if n == token {
			fmt.Println("Selamat Anda Berhasil Login")
			break
		}
	}
}
