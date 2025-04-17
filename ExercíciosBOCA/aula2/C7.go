package main

import (
	"fmt"
	"math"
)

func main() {

	var numero, valorTotal float64

	fmt.Scan(&numero)
	sinal := -1.0

	for i := 0.0; i < 20; i++ {
		if i == 0 {
			valorTotal += numero
		} else {
			valorTotal += (math.Pow(numero, float64(i)) / fatorialC7(i)) * sinal
			sinal *= -1
		}
	}
	fmt.Print(valorTotal)
}

func fatorialC7(n float64) float64 {
	resultado := 1.0
	for i := 2.0; i <= n; i++ {
		resultado *= float64(i)
	}
	return resultado
}