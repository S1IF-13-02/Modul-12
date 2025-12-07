package main


import "fmt"

func main() {
	var username, pass string
	var try int 
	
	fmt.Scan(&username, &pass)
	
	for username != "Admin" || pass != "Admin"{
		try++
		fmt.Scan(&username, &pass)
	}
	fmt.Println(try,"percobaan gagal login")
}