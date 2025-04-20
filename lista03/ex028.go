package main

import (
	f "fmt"
	m "math"
)

func main() {
	var resultado, produto float64

	contador := 1.0
	sinal := 1.0

	for i := 1.0; i <= 51; i++ {
		potencia := m.Pow(contador, 3)
		valorTotal := 1 / potencia
		resultado += sinal * valorTotal
		sinal *= -1
		contador += 2
	}

	produto = m.Pow((resultado * 32), 1.0/3.0)
	f.Printf("O valor do ∏ com 51 termos é igual a: %.15f", produto)
}
