package main

import (
	f "fmt"
	m "math"
)

func main() {
	var (
		Somatorio float64
		N         int
	)

	f.Scan(&N)
	vetor := make([]float64, N)
	for i := 0; i < N; i++ {
		f.Scan(&vetor[i])
	}

	for i := 0; i <= (N/2)-1; i++ {
		sub := vetor[i] - vetor[N-1-i]
		Somatorio += m.Pow(sub, 3)
	}
	f.Printf("%.2f", Somatorio)
}
