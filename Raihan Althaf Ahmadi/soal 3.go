package main

import "fmt"
func main (){
	var x, y, total int
	fmt.Print("Masukan 2 Bilangan Bulat Untuk di bagi : ")
	fmt.Scan(&x, &y)
	
	total = 0
	for x >= y {
		x = x - y
		total++
	}
	fmt.Print(total)
}