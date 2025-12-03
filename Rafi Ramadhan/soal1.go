ackage main

import (
	"fmt"
)

func main() {
	correctUser := "Admin"
	correctPass := "Admin"

	var user, pass string
	failed := 0

	for {
		fmt.Print("Masukkan username dan password: ")
		fmt.Scan(&user, &pass)

		if user == correctUser && pass == correctPass {
			break
		}

		failed++
	}

	fmt.Println(failed, "percobaan gagal login")
}
