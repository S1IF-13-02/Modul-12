package main 

import "fmt"

func main() {

	const correctUsername = "Admin" 
	const correctPassword = "Admin" 

	percobaanGagal := 0
	isLoggedIn := false
	
	var username, password string

	for !isLoggedIn{
		fmt.Print("Masukkan Username: ")
		fmt.Scan(&username)
		fmt.Print("Masukkan Password: ")
		fmt.Scan(&password)

		if username == correctUsername && password == correctPassword{
			isLoggedIn = true
			fmt.Println("Login berhasil!")
			fmt.Printf("Jumlah percobaan login yang gagal: %d", percobaanGagal)
		}else{
			percobaanGagal++
			fmt.Println("Username atau Password salah. Silahkan coba lagi!")
		}
	} 
}
