package main

import (
	"fmt"
	"math"
)



func main() {

	var horas,valor float64

	fmt.Println("Quantidade de horas que a charrete foi usada: ")
	fmt.Scan(&horas)

	if horas <= 3 {
		valor = horas * 5
	} 

	if horas > 3 {
		valor = math.Floor(float64(horas/3)) * 10 + (horas - (math.Floor(float64(horas/3)) * 3)) * 5
	}

	
	fmt.Printf("O valor a ser pago é R$ %.2f\n", valor)
}