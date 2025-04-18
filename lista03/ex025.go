package main

import (
	f "fmt"
	m "math"
)

func main() {
	var (
		quadrado = 15.0
		soma     = 0.0
		sinal    = 1.0
	)

	for i := 0.0; i <= 15; i++ {
		valorParcial := m.Pow(2, i) / m.Pow(quadrado, 2)
		soma += valorParcial * sinal
		sinal *= -1.0

		if quadrado == 1 {
			break
		} else {
			quadrado--
		}
	}
	f.Printf("%.2f", soma)
}
