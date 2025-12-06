package main

import "fmt"

func main() {
	correctuser := "Admin"
	correctpass := "Admin"
	var ipass, iuss string
	failed := 0

	for {
		fmt.Print("masukan ussername dan pass : ")
		fmt.Scan(&ipass, &iuss)

		if ipass == correctpass && iuss == correctuser {
			break
		}
		failed++
	}
	fmt.Println("percobaan gagal login", failed)
}
