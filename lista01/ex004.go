package main

import (
	"fmt"
)
func main() {

	var (
		salario,energia float64
		custo = [3]float64{}
	)
	
	fmt.Print("Qual o valor do salário mínimo? R$")
	fmt.Scan(&salario)
	fmt.Print("Qual a quantidade KW gasto? KW")
	fmt.Scan(&energia)

	custo[0] = (salario * 0.7) / 100
	custo[1] = custo[0] * energia
	custo[2] = custo[1] * 0.9

	fmt.Printf("Custo por kW: R$%.2f\n", custo[0])
	fmt.Printf("Custo do consumo: R$%.2f\n", custo[1])
	fmt.Printf("Custo com desconto: R$%.2f\n",  custo[2])

}

