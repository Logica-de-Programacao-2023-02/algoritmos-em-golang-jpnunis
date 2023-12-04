package main

import "fmt"

func main() {
    var salarioAtual, novoSalario float64

    fmt.Print("Digite o salário atual do funcionário: R$")
    fmt.Scan(&salarioAtual)

    // Verificar condições e calcular o novo salário
    if salarioAtual < 1000.0 {
        novoSalario
