package main

import "fmt"

func main() {
	var N, x, y, j, temp int
	fmt.Scan(&N)
	x = 0
	y = 1
	j = 0
	for j < N {
		fmt.Print(x, " ")
		temp = x + y
		x = y
		y = temp
		j = j + 1
	}
}
