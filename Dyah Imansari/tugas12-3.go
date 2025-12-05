package main
import "fmt"
func main() {
	var x, y, i int
	fmt.Print("Masukkan dua bilangan bulat positif (x y): ")
	fmt.Scan(&x, &y)
	if x < y || y <= 0 {
		fmt.Println("Input tidak valid. Pastikan x >= y dan y > 0.")
		return
	}
	i = 0
	for x >= y {
		x = x-y
		i = i+1
	}
	fmt.Println("hasil integer division:", i)
}