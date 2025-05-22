package main

import f "fmt"

func main() {
	var (
		N, M, X int
	)

	f.Scan(&N, &M)

	// Cria a matriz original e transposta com o tamanho que coloca
	matrix := make([][]int, N)
	transposta := make([][]int, M)
	for i := 0; i < N; i++ {
		matrix[i] = make([]int, M)
	}
	for j := 0; j < M; j++ {
		transposta[j] = make([]int, N)
	}

	// Armazena o valor
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			f.Scan(&X)
			matrix[i][j] = X
		}
	}

	// Inverter matriz
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			transposta[j][i] = matrix[i][j]
		}
	}

	// Exibe a matriz transposta corretamente
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			f.Printf("%d ", transposta[i][j])
		}
		if i != M-1 {
			f.Println()
		}
	}
}
