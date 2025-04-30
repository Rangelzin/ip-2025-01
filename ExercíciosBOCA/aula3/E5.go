package main

import (
	f "fmt"
)

func main() {
	var qtd int

	f.Scan(&qtd)
	vetor := make([]int, qtd)
	for i := 0; i < qtd; i++ {
		f.Scan(&vetor[i])
	}

	segCres := 1
	maxSeq := 1
	for i := 1; i < qtd; i++ {
		if vetor[i] > vetor[i-1] {
			segCres++
			if segCres > maxSeq {
				maxSeq = segCres
			}
		} else {
			segCres = 1
		}
	}
	f.Print(maxSeq)
}
