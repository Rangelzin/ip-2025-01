package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um valor inteiro para verificação: ")
	fmt.Scan(&numero)

	if numero > 0 {
		fmt.Println("O valor informado é positivo")
	} else if numero < 0 {
		fmt.Println("O valor informado é negativo")
	} else {
		fmt.Println("O valor informado é nulo")
	}
}