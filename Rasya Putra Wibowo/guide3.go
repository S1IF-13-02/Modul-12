package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan N: ")
    fmt.Scan(&n)

    a,b := 0,1
    i := 0

	for i < n {
		fmt.Print(a, " ")
		a,b = b, a+b 
		i++
	}
}