package main

import f "fmt"

func main() {
	var (
		N, M, K, T int
	)
	f.Scan(&N, &K, &M)

	// Cria a matriz 1 com o tamanho que coloca
	matrizA := make([][]int, N)
	matrizB := make([][]int, K)
	matrizC := make([][]int, N)

	montarMatriz(&matrizA, N, K)
	montarMatriz(&matrizB, K, M)
	montarMatriz(&matrizC, N, M)

	guardarValor(&matrizA, N, K, T)
	guardarValor(&matrizB, K, M, T)

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			soma := 0
			for k := 0; k < K; k++ {
				soma += matrizA[i][k] * matrizB[k][j]
			}
			matrizC[i][j] = soma
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			f.Printf("%d ", matrizC[i][j])
		}
		f.Println()
	}

}

func montarMatriz(l *[][]int, X, Y int) {
	for i := 0; i < X; i++ {
		(*l)[i] = make([]int, Y)
	}
}

func guardarValor(l *[][]int, X, Y, Z int) {
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			f.Scan(&Z)
			(*l)[i][j] = Z
		}
	}
}
