package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan angka: ")
    fmt.Scan(&n)

	if n == 0{
		fmt.Print(1)
	}else{
		for i := n; i >= 1; i--{
			fmt.Print(i)
			if i != 1 {
				fmt.Print("x")
			}
		}
	}
}