package main

import (
	"fmt"
)

func main() {

	var (
		num int
		soma float32
	)

	fmt.Print("Digite um numero para somar suas contas decimais: ")
	fmt.Scan(&num)

	if num <= 1 {
		fmt.Print("Numero invalido!")
	} else {
		for i := 1; i <= num; i++ {
			soma = soma + (1.0 / float32(i)) 
		}
	}

	fmt.Printf("O Valor do somatório: %.6f\n", soma)

}