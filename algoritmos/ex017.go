package main

import (
	"fmt"
)

func main() {

	var (
		valInit, qtd int
	)

	fmt.Print("Insira o valor inicial / Quantidade de valores: ")
	fmt.Scan(&valInit, &qtd)


	if valInit % 2 != 0 {
		fmt.Print("O primeiro numero não é par!\n")
	} else {
		for i := 0; i < qtd; i++ {
			fmt.Printf("%d\n",valInit)
				valInit += 2
		}
	}
}