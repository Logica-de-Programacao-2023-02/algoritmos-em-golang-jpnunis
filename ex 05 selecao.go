package main

import "fmt"

func main() {
    var numero int
    fmt.Print("Digite um número inteiro: ")
    fmt.Scan(&numero)

    if numero%3 == 0 && numero%5 == 0 {
        fmt.Println("Múltiplo de 3 e de 5.")
    } else {
        fmt.Println("Não é múltiplo de 3 e de 5 simultaneamente.")
    }
}
