package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n == 0 {
		fmt.Println(1)
	} else {
		for i := n; i >= 1; i-- {
			fmt.Print(i)
		if i != 1 {
			fmt.Print("x")
		}
}
	}	
}
	