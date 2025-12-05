package main
import "fmt"
func main() {
	var username, password string
	var total int
	total = 0
	for {
	fmt.Println("Masukkan Username dan Password: ")
	fmt.Scan(&username)
	fmt.Scanln(&password)
	if username == "Admin" && password == "Admin" {
		break
	}
	total++
	}
	fmt.Printf("%d percobaan gagal login\n", total)
}