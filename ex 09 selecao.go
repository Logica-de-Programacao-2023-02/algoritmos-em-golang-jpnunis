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

	numeros := []float64{num1, num2, num3}
	sort.Float64s(numeros)

	fmt.Println("Números em ordem crescente:", numeros)
}
