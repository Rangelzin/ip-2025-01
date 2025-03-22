package main 

import "fmt"

var (
	conta,consumo int
	valor float32
	tipo string
)

func main() {
	fmt.Print("Digite sua CONTA / CONSUMO / TIPO: ")
	fmt.Scan(&conta, &consumo, &tipo)

	if tipo = "R" || tipo = "r" {

		valor = (0.05 * consumo) + 5

		fmt.Print("\nCONTA = " + conta)
		fmt.Print("VALOR DA CONTA = R$" + valor)

	} else if tipo = "C" || tipo = "c"{

		if consumo > {
			
		}

		fmt.Print("\nCONTA = " + conta)
		fmt.Print("VALOR DA CONTA = R$" + valor)

	} else if tipo = "I" || tipo = "i" {

		if consumo > {
			
		}

		fmt.Print("\nCONTA = " + conta)
		fmt.Print("VALOR DA CONTA = R$" + valor)

	} else {
		fmt.Print("TIPO INVÁLIDO!")
	}
}