package main
import "fmt"
func main() {
	var x,y int
	fmt.Print("Masukkan x dan y : ")
	fmt.Scan(&x,&y)
	j := 0
	for x >= y {
		x = x - y
		j++
	}
	fmt.Print("Output : ", j)
}