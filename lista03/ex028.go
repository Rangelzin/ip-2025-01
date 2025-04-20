package main

import (
	f "fmt"
	m "math"
)

func main() {
	var (
		resultado float64
		sinal     float64 = 1
		produto   float64
		contador  float64 = 1
	)

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
