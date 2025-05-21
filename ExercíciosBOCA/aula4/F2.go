package main

import f "fmt"

func main() {
	var (
		N, M, X int
	)

	f.Scan(&N, &M)

	sumR := make([]int, N)
	sumC := make([]int, M)

	// Cria a matriz original com o tamanho que colocou
	matrix := make([][]int, N)
	for i := 0; i < N; i++ {
		matrix[i] = make([]int, M)
	}

	// Armazena o valor e ja realiza as somas
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			f.Scan(&X)
			sumC[j] += X
			sumR[i] += X
			matrix[i][j] = X
		}
	}

	// Exibe as somas
	for i := 0; i < N; i++ {
		f.Printf("%d ", sumR[i])
		if i == N-1 {
			f.Println()
		}
	}
	for i := 0; i < M; i++ {
		f.Printf("%d ", sumC[i])
	}
}
