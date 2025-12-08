package main
import "fmt"
func main(){
	var n int
	fmt.Print("Masukan bilangan Bulat non Negatif :")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("1")
		return
	}
		i := n
		for {
			fmt.Print(i)
			if i > 1 {
				fmt.Print(" X ")
			}
			i--
			if i == 0 {
				break
			}
	}

}