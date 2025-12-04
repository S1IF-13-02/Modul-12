package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukan Angka : ")
	fmt.Scan(&n)
	
	if n < 0 {
        fmt.Println("Input harus bilangan bulat non negatif")
        return
    }

	i := n

	for i > 1 {
		 fmt.Printf("%d x ", i)
		 i = i - 1 
	}
	fmt.Print(1)
}