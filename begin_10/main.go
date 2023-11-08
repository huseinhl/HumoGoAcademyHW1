package main

import "fmt"
import "math"

func main(){
	var a int
	var b int
	fmt.Println("задать числа")

	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println(a+b)
	fmt.Println(a*b)
	fmt.Println(a/b)
	fmt.Println(math.Sqrt(float64(a*b)))
}