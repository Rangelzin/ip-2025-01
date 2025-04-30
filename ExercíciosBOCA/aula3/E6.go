package main

import (
	f "fmt"
	s "sort"
)

func main() {
	var qtd int
	f.Scan(&qtd)

	vetor := make([]int, qtd)
	for i := 0; i < qtd; i++ {
		f.Scan(&vetor[i])
	}
	s.Ints(vetor)
	for i := 0; i < qtd; i++ {
		f.Printf("%d ", vetor[i])
	}
}
