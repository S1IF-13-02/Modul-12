package main 

import "fmt"

func main(){
	var user string 
	var pass string 
	fmt.Scan(&user, &pass)
	percobaan := 0
	for user != "Admin" || pass != "Admin" {
		fmt.Scan(&user, &pass)
		percobaan = percobaan + 1
	}
	fmt.Println(percobaan,"percobaan gagal login")
}