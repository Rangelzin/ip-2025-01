package main 

import "fmt"

func main() {

	var numero int

	fmt.Print("Digite um valor inteiro para verificação: ")
	fmt.Scan(&numero)

	if numero % 2 == 0 {
		fmt.Println("O número informado é um valor par")
	} else {
		fmt.Println("O número informado é um valor ímpar")
	}
}