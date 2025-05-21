package main

import f "fmt"

func main() {
	var (
		N, M, X int
	)
	f.Scan(&N, &M)

	// Cria a matriz 1 e 2 com o tamanho que coloca
	matrix1 := make([][]int, N)
	matrix2 := make([][]int, N)
	for i := 0; i < N; i++ {
		matrix1[i] = make([]int, M)
		matrix2[i] = make([]int, M)
	}

	// Armazena o valor da matriz1
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			f.Scan(&X)
			matrix1[i][j] = X
		}
	}
	
}
