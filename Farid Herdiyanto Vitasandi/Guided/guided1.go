package main

import "fmt"

func main() {

	var bilangan int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan < 0 {
		fmt.Print("Bilangan tidak boleh negatif")
		return
	}

	if bilangan == 0 {
		fmt.Println(1)
	}else{
		for i := bilangan; i >= 1; i--{
			fmt.Print(i)
			if i != 1 {
				fmt.Print(" x ")
			}
		}
	}
}

