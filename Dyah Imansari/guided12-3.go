package main
import "fmt"
func main() {
	var N, a, b, j, temp int
	fmt.Scan(&N)
	if N < 2 {
		fmt.Println("Masukkan minimal 2")
		return
	}
	a = 0
	b = 1
	j = 0
	for j < N {
		fmt.Print(a," ")
		temp = a + b
		a = b
		b = temp
		j++
	}
}