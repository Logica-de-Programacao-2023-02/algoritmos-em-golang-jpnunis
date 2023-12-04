package main

import (
	"fmt"
)

func main() {
	var (
		numero, soma int
		contador     float64
	)

	for {
		fmt.Print("Digite um número inteiro (ou 0 para encerrar): ")
		fmt.Scan(&numero)

		if numero == 0 {
			break
		}

		soma += numero
		contador++
	}

	if contador > 0 {
		media := float64(soma) / contador
		fmt.Printf("Média aritmética: %.2f\n", media)
	} else {
		fmt.Println("Nenhum número fornecido para calcular a média.")
	}
}
