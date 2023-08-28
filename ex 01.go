package main

import "fmt"

func main() {
	x := 21
	y := 23
	if x < y {
		fmt.Println("y é maior que x")
	} else if x > y {
		fmt.Println("x é maior que y")
	} else {
		fmt.Println("x é igual a y")
	}
}
