package main

import (
	"fmt"
    "math"
)
func main(){
	var a int 

	fmt.Println("задать числа")

	fmt.Scan(&a)

	var b int

	fmt.Scan(&b)

	fmt.Println(math.Sqrt(float64(a*b)))


}