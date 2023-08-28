package main

import "fmt"

func main() {
	i := 14
	if i%2 == 0 {
		fmt.Println("i é par")
	} else if i%2 != 0 {
		fmt.Println("i é impar")
	}

}
