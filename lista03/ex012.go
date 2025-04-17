package main

import (
	f"fmt"
)

func main() {

	var numero float64

	f.Print("Digite um número: ")
	f.Scan(&numero)

	resultado := 0.0
	sinal := 1.0

	for n := 0.0; n < 20; n++ {
		resultadoFat := 1.0

		if n == 0 {
			resultadoFat = 1.0
		} else {
			for i := n; i >= 2; i-- {
				resultadoFat *= i
			}
		}

		valorTotal := numero / resultadoFat
		resultado += sinal * valorTotal
		sinal *= -1
	}
	f.Printf("O resultado do somatório: %.5f\n", resultado)
}

