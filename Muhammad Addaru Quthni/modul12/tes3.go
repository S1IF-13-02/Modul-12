package main
import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&n)
	
	fmt.Println("Deret Fibonacci:")
	i := 0
	a, b := 0, 1
	for i < n {
		fmt.Print(a, " ")
		a, b = b, a+b
		i++
	}
	fmt.Println()
}