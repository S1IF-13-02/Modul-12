package main
import "fmt"

func main(){
	var n int

	fmt.Print("masukan bilangan:")
	fmt.Scan(&n)

	if n == 0{
		fmt.Println("1")
		return
	}
	i := n
	for i > 0 {
		fmt.Print(i)
		if i > 1 {
			fmt.Print("x")
		}
		i-- 
	}
	fmt.Println()
}
