package main

import "fmt"

func main() {
	var x, y int
	var ok bool = false

	fmt.Print("Masukkan x: ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan y: ")
	fmt.Scanln(&y)

	tmp := 0
	div := 0
	for !ok {
		if tmp+y <= x {
			tmp += y
			div++
		} else {
			break
		}
	}

	fmt.Println(div)
}
