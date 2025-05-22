package main

import f "fmt"

func main() {
	var (
		N, X int
	)
	f.Scan(&N)

	// Cria a matriz 1 com o tamanho que coloca
	matriz1 := make([][]int, N)
	matriz2 := make([][]int, N)
	for i := 0; i < N; i++ {
		matriz1[i] = make([]int, N)
		matriz2[i] = make([]int, N)
	}
	// Armazena o valor da matriz1
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			f.Scan(&X)
			matriz1[i][j] = X
		}
	}

	for j := 0; j < N; j++ {
		for i := N - 1; i >= 0; i-- {
			matriz2[j][i] = matriz1[i][j]
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			f.Printf("%d ", matriz2[i][j])
		}
		f.Println()
	}
}
