package main

import f "fmt"

func main() {
	var (
		vetOrigin = []float64{}
		qtd       int
	)

	f.Scan(&qtd)
	vetOrigin = make([]float64, qtd)
	for i := 0; i < qtd; i++ {
		f.Scan(&vetOrigin[i])
	}

	if len(vetOrigin) == 1 {
		f.Print(0)
		return
	}

	newVet := make([]float64, len(vetOrigin))
	for i := 0; i < qtd; i++ {
		if i == 0 {
			newVet[i] = vetOrigin[i+1]
		} else if i == qtd-1 {
			newVet[i] = vetOrigin[qtd-2]
		} else {
			newVet[i] = vetOrigin[i-1] + vetOrigin[i+1]
		}
	}
	for i := 0; i < qtd; i++ {
		f.Printf("%.0f ", newVet[i])
	}
}
