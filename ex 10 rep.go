package main

import "fmt"

func main() {
	var (
		numero, maior int
		leitura       int
	)

	// Ler o primeiro número
	fmt.Print("Digite um número inteiro (ou 0 para encerrar): ")
	fmt.Scan(&leitura)
	maior = leitura

	// Verifica
