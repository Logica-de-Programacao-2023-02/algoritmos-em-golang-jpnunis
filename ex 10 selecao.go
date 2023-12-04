package main

import "fmt"

func main() {
	var idade int

	fmt.Print("Digite a idade da pessoa: ")
	fmt.Scan(&idade)

	switch {
	case idade <= 9:
		fmt.Println("Mirim")
	case idade >= 10 && idade <= 13:
		fmt.Println("Infantil")
	case idade >= 14 && idade <= 17:
		fmt.Println("Juvenil")
	default:
		fmt.Println("Adulto")
	}
}
