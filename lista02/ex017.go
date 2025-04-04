package main

import "fmt"

func main() {

	var (
		conta int
		consumo, valor float32
		tipo string
	)

	fmt.Print("Digite a conta / tipo de conta / consumo: ")
	fmt.Scan(&conta, &tipo, &consumo)

	if tipo == "r" || tipo == "R"{
		valor = (consumo * 0.05) + 5
	} else if tipo == "c" || tipo == "C"{

		if consumo <= 80 {
			valor = 500
		
		} else {
			valor = ((consumo - 80) * 0.25) + 500
		}
	} else if tipo == "i" || tipo == "I"{
		if consumo <= 100 {
			valor = 800
		} else {
			valor = ((consumo - 100) * 0.04) + 800
		}
	}
		fmt.Printf("Conta = %d\n", conta)
		fmt.Printf("Valor da Conta = %.2f\n", valor)
}