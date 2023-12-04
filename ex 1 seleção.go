package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Digite o primeiro número inteiro: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número inteiro: ")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Printf("O maior número é: %d\n", num1)
	} else {
		fmt.Printf("O maior número é: %d\n", num2)
	}
}
