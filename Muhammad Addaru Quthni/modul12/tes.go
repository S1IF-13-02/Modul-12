package main
import "fmt"
func main() {
	var i int
	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&i)
	if i == 0 {
		fmt.Println("1")
	} else {
		for i >= 1 {
			fmt.Print(i)
			if i > 1 {
				fmt.Print(" x ")
			}
			i--
		}
		fmt.Println()
	}
}