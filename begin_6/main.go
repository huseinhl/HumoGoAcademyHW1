package main

import "fmt"

func main(){
	var a int

	fmt.Println("длины ребер")

	fmt.Scan(&a)

	var b int

	fmt.Scan(&b)

	var c int 

	fmt.Scan(&c)

	fmt.Println("объем:")
    fmt.Println(a*b*c)

	fmt.Println("площадь:")
	fmt.Println(2*(a*b+b*c+a*c))
}