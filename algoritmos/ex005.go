package main

import (
	"fmt"
)
var (
	conta int
	valor,consumo float32
	tipo string
)

func main() {
	fmt.Print("Digite sua CONTA / CONSUMO / TIPO: ")
	fmt.Scan(&conta, &consumo, &tipo)

	if tipo == "R" || tipo == "r" {
		valor = (0.05 * consumo) + 5
	} else if tipo == "C" || tipo == "c"{
		if consumo <= 80 {
			valor = 500
		} else {
			valor = 500 + (0.25 * (consumo - 80))
		}
	} else if tipo == "I" || tipo == "i" {
		if consumo <= 100 {
			valor = 800
		} else {
			valor = 800 + (0.04 * (consumo - 100))
		}
	} else {
		fmt.Print("TIPO INVÁLIDO!")
	}

	fmt.Println("CONTA = R$", conta)
	fmt.Println("VALOR DA CONTA = R$", valor)
}