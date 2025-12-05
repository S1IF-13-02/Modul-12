package main

import "fmt"

func main() {
	var a int
	var j int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&a)
	j = a
	for j > 1 {
		fmt.Print(j, " x ")
		j = j - 1
	}
	fmt.Println(1)
}
