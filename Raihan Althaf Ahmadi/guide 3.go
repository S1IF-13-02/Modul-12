package main

import "fmt"
func main (){
	var n int
	fmt.Print("Masukan Bilangan Bulat : ")
	fmt.Scan(&n)
	a, b := 0, 1
	i := 0
	for i < n {
		fmt.Print(a, " ")
		total := a+b
		a = b
		b = total
		n--
	}
}