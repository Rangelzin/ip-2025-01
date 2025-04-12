package main

import (
	"fmt"
	"math"
)

func main() {

	var numero float64

	fmt.Scan(&numero)

	resultado := 0.0
	sinal := 1.0

	for n := 0; n < 20; n++ {
		valorTotal := math.Pow(numero, float64(n)) / factorial(n)
		resultado += sinal * valorTotal
		sinal *= -1
	}
	fmt.Print(resultado)
}

func factorial(n int) float64 {
	resultado := 1.0
	for i := n; i >= 2; i-- {
		resultado *= float64(i)
	}
	return resultado
}
