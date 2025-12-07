package main

import (
	"fmt"
)

func main() {
	var bilangan int64

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&bilangan)

	if bilangan < 0 {
		fmt.Println("Input harus bilangan bulat positif.")
	}

	for bilangan > 0 {

		digitTerakhir := bilangan % 10

		fmt.Println(digitTerakhir)

		bilangan = bilangan / 10
	}
}
