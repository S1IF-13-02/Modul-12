package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("masukan nilai x dan y : ")
	fmt.Scan(&x, &y)

	j := 0

	for x >= y {
		x = x - y
		j++
	}
	fmt.Println("keluaran: ", j)

}
