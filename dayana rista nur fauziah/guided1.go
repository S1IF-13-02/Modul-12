package main

import "fmt"

func main() {
	var bilangan, j int

	fmt.Print("masykan bilangan bulat non negatif (n):")
	fmt.Scan(&bilangan)

	j = bilangan
	for j > 1 {
		fmt.Print(j, "x")
		j = j - 1
	}
	fmt.Println(1)
}
