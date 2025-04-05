package main

import "fmt"

func main() {

	var vendasBruta, comissao, salario float64

	fmt.Scan(&vendasBruta)

	comissao = vendasBruta * 0.09

	salario = 500 + comissao

	if vendasBruta > 15000 {salario += 800}
	fmt.Printf("%.5f\n", salario)
}