package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukan n: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println(1)
		return
	}

	for n > 0 {
		fmt.Print(n, " ")

		if n != 1 {
			fmt.Print("x" + " ")
		}

		n--
	}
	
	fmt.Println()
}
