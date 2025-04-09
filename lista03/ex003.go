package main

import (
	"fmt"
)

func main() {
	var salarioCarlos, salarioJoao, MontCarlos, MontJoao float64
	var tempo int

	fmt.Print("Digite o valor do salário do Carlos: ")
	fmt.Scan(&salarioCarlos)

	salarioJoao = salarioCarlos/3
	MontCarlos, MontJoao = salarioCarlos, salarioJoao

	for MontJoao < MontCarlos {
		MontCarlos = MontCarlos * 1.02
		MontJoao = MontJoao * 1.05
		tempo += 1
	}
	fmt.Printf("Tempo de meses necessários: %d\n",tempo)
}