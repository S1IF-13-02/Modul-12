package main
import "fmt"
func main() {
	var n, digit int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Input harus bilangan bulat positif")
		return
	}
	for n > 0 {
		digit = n % 10
		fmt.Println(digit)
		n = n/10
	}
}