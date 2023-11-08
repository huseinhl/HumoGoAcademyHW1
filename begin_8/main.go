package main

import "fmt"

func main(){
	var a int

	fmt.Println("задать числа")

	fmt.Scan(&a)

	var b int

	fmt.Scan(&b)
    
	fmt.Println("среднее арифметичское:")
	fmt.Println((a+b)/2)
}