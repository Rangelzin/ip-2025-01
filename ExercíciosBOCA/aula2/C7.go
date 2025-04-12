package main

import (
	"fmt"
	"math"
)

func main() {

	var numero, valorTotal float64

	fmt.Scan(&numero)

	for i := 0; i < 20; i++ {
		fatorial := 1
		if i == 0 {
			valorTotal += numero
		} else if i%2 != 0 {
			for j := i; j >= 1; j-- {
				fatorial *= j
			}
			valorTotal -= math.Pow(numero, float64(i)) / float64(fatorial)
		} else {
			for j := i; j >= 1; j-- {
				fatorial *= j
			}
			valorTotal += math.Pow(numero, float64(i)) / float64(fatorial)
		}
	}
	fmt.Print(valorTotal)
}
