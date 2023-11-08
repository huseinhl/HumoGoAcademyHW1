package main 

import "fmt"

func main(){
	var a int

	fmt.Println("сторона a:")

	fmt.Scan(&a)

	var b int

	fmt.Println("сторона b:")

    fmt.Scan(&b)

	fmt.Println((a + b) * 2)

	fmt.Println(a * b)
}