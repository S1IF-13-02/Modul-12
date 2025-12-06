package main

import (
	"fmt"
)

func main() {

	var x, y int
	hasilbagi := 0

	fmt.Print("masukan x dan y : ")
	fmt.Scan(&x, &y)

	if y == 0 {
		fmt.Print("input y tidak boleh 0")
		return
	}

	for x >= y {
		x = x - y
		hasilbagi = hasilbagi + 1

	}
	fmt.Print(hasilbagi)
}
