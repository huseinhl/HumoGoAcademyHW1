package main
 
import "fmt"

func main(){
	var a int

	fmt.Println("длина ребра")

	fmt.Scan(&a)
	fmt.Println("объем:")

	fmt.Println(a*a*a)
	fmt.Println("площадь:")
	fmt.Println(6 * a * a)
}