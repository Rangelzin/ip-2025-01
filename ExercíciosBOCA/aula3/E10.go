package main

import f "fmt"

func main() {
	var qtd, teste = 0, true
	f.Scan(&qtd)

	vetor := make([]int, qtd)
	for i := 0; i < qtd; i++ {
		f.Scan(&vetor[i])
	}
	for i := 0; i < qtd; i++ {
		if vetor[i] == vetor[qtd-i-1] {
			teste = true
		} else {
			teste = false
		}
	}

	if teste {
		f.Printf("SIM")
	} else {
		f.Printf("NAO")
	}
}
