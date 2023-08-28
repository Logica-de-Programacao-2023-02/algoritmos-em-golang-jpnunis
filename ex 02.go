package main

import "fmt"

func main() {
	x := 6
	y := 13
	z := 11
	if x < y && x < z {
		fmt.Println("x é menor que y ")
	} else if y < x && y < z {
		fmt.Println("y é menor que z ")
	} else if z < y && z < x {
		fmt.Println("z é menor que x ")
	}
}
