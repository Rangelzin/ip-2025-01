package main

import f "fmt"

func main() {
	var qtd int
	f.Scan(&qtd)

	vetor := make([]int, qtd)
	conjunto := make(map[int]bool)
	for i := 0; i < qtd; i++ {
		f.Scan(&vetor[i])
		conjunto[vetor[i]] = true
	}
	f.Printf("%d ", len(conjunto))
}
