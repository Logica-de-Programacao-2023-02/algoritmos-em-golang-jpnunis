package main

import "fmt"

func main() {
	var altura float64
	var sexo string

	fmt.Print("Digite a altura (em metros): ")
	fmt.Scan(&altura)

	fmt.Print("Digite o sexo (M para masculino, F para feminino): ")
	fmt.Scan(&sexo)

	pesoIdeal := 0.0

	switch sexo {
	case "M":
		pesoIdeal = (72.7 * altura) - 58
	case "F":
		pesoIdeal = (62.1 * altura) - 44.7
	default:
		fmt.Println("Sexo inválido. Use 'M' para masculino ou 'F' para feminino.")
		return
	}

	fmt.Print("Digite o peso da pessoa: ")
	var pesoAtual float64
	fmt.Scan(&pesoAtual)

	if pesoAtual < pesoIdeal {
		fmt.Println("A pessoa está abaixo do peso ideal.")
	} else if pesoAtual > pesoIdeal {
		fmt.Println("A pessoa está acima do peso ideal.")
	} else {
		fmt.Println("A pessoa está dentro do peso ideal.")
	}
}
