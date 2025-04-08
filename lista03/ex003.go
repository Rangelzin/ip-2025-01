package main

import (
	"fmt"
)

func main()  {
	var salarioC, salarioJ, MontC, MontJ float64
	var tempo int

	fmt.Print("Digite o valor do salário do Carlos: ")
	fmt.Scan(&salarioC)

	salarioJ = salarioC/3
	MontC, MontJ = salarioC, salarioJ

	for MontJ < MontC {
		MontC = MontC * 1.02
		MontJ = MontJ * 1.05
		tempo += 1
	}
	fmt.Printf("Tempo de meses necessários: %d\n",tempo)
}