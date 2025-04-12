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
		valorTotal := math.Pow(numero, float64(2*n)) / fatorial2(2*n)
		resultado += sinal * valorTotal
		sinal *= -1
	}

	cosNUM := math.Cos(numero)
	dif := resultado - cosNUM

	fmt.Printf("%.4f %.4f %.4f", resultado, cosNUM, dif)
}

func fatorial2(n int) float64 {
	resultado := 1.0
	for i := n; i >= 2; i-- {
		resultado *= float64(i)
	}
	return resultado
}
