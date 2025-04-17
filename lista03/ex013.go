package main

import f "fmt"

func main() {
	var H, soma, numerador, denominador float64
	for i := 1.0; i <= 50; i++ {
		numerador = (i * 2) - 1
		denominador = i
		H = numerador / denominador
		soma += H
	}
	f.Printf("Soma = %.2f\n", soma)
}
