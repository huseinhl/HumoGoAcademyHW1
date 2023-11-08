package main

import "fmt"

func main(){
	var d int

	fmt.Println("диаметр:")

	fmt.Scan(&d)
    
    fmt.Println(float64(d * 3.14))
}