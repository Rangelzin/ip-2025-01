package main

import (
	"fmt"
	"math"
)

func main() {
	var numero float64

	for {
		fmt.Print("Digite um numero: ")
		fmt.Scan(&numero)
		encontrado := false

		if numero > 0 {
			for i := 1.0; math.Pow(i, 3) <= numero; i++ {

				forms := i * (i + 1) * (i + 2)

				if forms == numero {
					encontrado = true
					break
				}
			}

			if encontrado {
				fmt.Print("O número é triangular\n")
			} else {
				fmt.Print("O número NÃO é triangular\n")

			}
		} else {
			return
		}
	}
}
