package main 

import "fmt"

func main(){
	var n int 
	fmt.Scan(&n)
	a := 0
	b := 1
	i := 0
	for i < n {
		fmt.Print(a, " ")
		c := a + b
		a = b
		b = c
		i = i + 1
	}
}