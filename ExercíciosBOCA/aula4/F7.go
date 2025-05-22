package main

import f "fmt"

func main() {
	var N, M, T int

	f.Scan(&N, &M)

	matrizA := make([][]int, N)
	montarMatriz2(&matrizA, N, M)

	guardarValor2(&matrizA, N, M, T)

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			big := 0.0
			for k := 0; k < M; k++ {
				if matrizA[i][k] > int(big) {
					big = float64(matrizA[i][k])
				}
			}
			f.Printf("%.6f ", float64(matrizA[i][j])/big)
		}
		f.Println()
	}

}

func montarMatriz2(l *[][]int, X, Y int) {
	for i := 0; i < X; i++ {
		(*l)[i] = make([]int, Y)
	}
}

func guardarValor2(l *[][]int, X, Y, Z int) {
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			f.Scan(&Z)
			(*l)[i][j] = Z
		}
	}
}
