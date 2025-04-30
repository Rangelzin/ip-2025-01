package main

import f "fmt"

func main() {
	var qtd, times int

	f.Scan(&qtd)
	vetorA := make([]int, qtd)
	vetorB := make([]int, qtd)

	ContagemDeDados(&vetorA, qtd)
	ContagemDeDados(&vetorB, qtd)

	for i := 0; i < qtd; i++ {
		times += (vetorA[i] * vetorB[i])
	}
	f.Printf("%d ", times)
}

func ContagemDeDados(lista *[]int, indice int) {
	for i := 0; i < indice; i++ {
		f.Scan(&(*lista)[i])
	}
}
