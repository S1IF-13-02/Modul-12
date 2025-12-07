package main

import "fmt"
func main (){
	var n, total int
	fmt.Print("Masukan Bilangan Bulat : ")
	fmt.Scan(&n)

	for n > 0{
		total = n%10
		fmt.Println(total)
		n /= 10
	}
}