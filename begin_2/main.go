package main

import "fmt"

func main(){
	var a1 int 

	fmt.Println("первая сторона квадрата")

	fmt.Scan(&a1)

	var a2 int
	
	fmt.Println("вторая сторона квадрата")

	fmt.Scan(&a2)

	fmt.Println(a1 * a2)
}