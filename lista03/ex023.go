package main

import f "fmt"

func main() {
	var (
		N    int
		soma float64 = 1000
	)

	f.Print("Digite a quantidade de termos: ")
	f.Scan(&N)

	subtração := soma
	sinal := -1.0
	for i := 2; i <= N; i++ {
		subtração -= 3
		valorNovo := (subtração / float64(i)) * sinal
		soma += valorNovo
		sinal *= -1.0
	}
	f.Printf("%.2f", soma)
}
