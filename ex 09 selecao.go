package main

import (
	"fmt"
	"sort"
)

func main() {
	var num1, num2, num3 float64

	fmt.Print("Digite o primeiro número real: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número real: ")
	fmt.Scan(&num2)

	fmt.Print("Digite o terceiro número real: ")
	fmt.Scan(&num3)

	// Criar uma fatia (slice) com os números e ordená-la
	numeros := []float64{num1, num2, num3}
	sort.Float64s(numeros)

	// Mostrar os números em ordem crescente
	fmt.Println("Números em ordem crescente:", numeros)
}
