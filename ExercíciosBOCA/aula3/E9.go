package main

import f "fmt"

func main() {
	var qtd, soma int
	f.Scan(&qtd)

	vetor := make([]int, qtd)
	newVetor := make([]int, qtd)
	for i := 0; i < qtd; i++ {
		f.Scan(&vetor[i])
		soma += vetor[i]
		newVetor[i] = soma
	}

	for _, v := range newVetor {
		f.Printf("%d ", v)
	}
}
