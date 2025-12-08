package main

import "fmt"

func main() {
	var num int
	var ok bool = false

	fmt.Print("Masukkan: ")
	fmt.Scanln(&num)

	for !ok {
		tmp := num % 10
		fmt.Println(tmp)
		num /= 10
		if num == 0 {
			ok = true
		}
	}
}
