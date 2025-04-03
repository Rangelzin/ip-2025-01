package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Digite um valor inteiro: ")
	fmt.Scan(&numero)


	if numero > 20 && numero < 90 {
		fmt.Println("Número informado pelo usuário ESTÁ compreendido entre 20 e 90")
	} else {
		fmt.Println("Número informado pelo usuário NÃO ESTÁ compreendido entre 20 e 90")
	}
}