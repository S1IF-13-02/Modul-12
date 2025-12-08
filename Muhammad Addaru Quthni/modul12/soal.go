package main
import "fmt"
func main() {
	var usn, pass string = "Admin", "Admin"
	var i int = 0
	fmt.Print("Masukkan Username dan Password: ")
	fmt.Scan(&usn, &pass)
	if usn != "Admin" || pass != "Admin" {
		for usn != "Admin" || pass != "Admin" {
			i++
			fmt.Scan(&usn, &pass)
		}
	}
	fmt.Print(i, " Percobaan gagal login ")
}