package main

import (
	"fmt"
)
func main() {

	var (
		salario,energia float64
		custo = [3]float64{1,2,3}
	)
	
	fmt.Print("Qual o valor do salário mínimo? R$")
	fmt.Scan(&salario)
	fmt.Print("Qual a quantidade KW gasto? KW")
	fmt.Scan(&energia)

	custo[0] = (salario * 0.7) / 100
	custo[1] = custo[0] * energia
	custo[2] = custo[1] * 0.9

	fmt.Println("Custo por kW: R$ ", custo[0])
	fmt.Println("Custo do consumo: R$ ", custo[1])
	fmt.Println("Custo com desconto: R$ ", + custo[2])

}

