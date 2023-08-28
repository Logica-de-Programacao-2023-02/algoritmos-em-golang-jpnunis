package main

import "fmt"

func main() {
	peso := 90.
	altura := 1.70
	imc := peso / (altura * altura)
	if imc < 18.5 {
		fmt.Println("peso normal")
	} else if imc >= 18.5 && imc < 24.9 {
		fmt.Println("peso normal")
	} else if imc >= 2 && imc < 29.9 {

	}
}