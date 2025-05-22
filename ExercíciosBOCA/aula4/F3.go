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

	// matriz de arrays possíveis vizinhos
	vizinhoPossíveis := [8][2]int {
		{-1,-1},{-1, 0},{-1, 1},{ 0,-1},{ 0, 1},{ 1,-1},{ 1, 0},{ 1, 1},
	}
	
	// Soma
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			soma := 0
			for _,d := range vizinhoPossíveis {
				ni := i + d[0]
				nj := j + d[1]
				if ni >= 0 && ni < N && nj >= 0 && nj < M {
					soma += matrix1[ni][nj]
				}
			}
			matrix2[i][j] = soma
		}
	}
	
	// 
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			f.Printf("%d ", matrix2[i][j])
		}
		f.Println()
	}
}


