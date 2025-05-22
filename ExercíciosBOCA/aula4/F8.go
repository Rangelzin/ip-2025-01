package main

import (
	f "fmt"
	"math"
)

func main() {
	var N, M int
	f.Scan(&N, &M)

	matrizA := make([][]int, N)
	montarMatriz3(&matrizA, N, M)
	guardarValor3(&matrizA, N, M)

	for i := 0; i < N; i++ {
		linha := matrizA[i]
		maxVal := linha[0]
		for j := 1; j < M; j++ {
			if linha[j] > maxVal {
				maxVal = linha[j]
			}
		}

		expVals := make([]float64, M)
		sumExp := 0.0

		for j := 0; j < M; j++ {
			expVals[j] = math.Exp(float64(linha[j] - maxVal)) // prevenção contra overflow
			sumExp += expVals[j]
		}

		for j := 0; j < M; j++ {
			soft := expVals[j] / sumExp
			f.Printf("%.6f ", soft)
		}
		f.Println()
	}
}

func montarMatriz3(l *[][]int, X, Y int) {
	for i := 0; i < X; i++ {
		(*l)[i] = make([]int, Y)
	}
}

func guardarValor3(l *[][]int, X, Y int) {
	var temp int
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			f.Scan(&temp)
			(*l)[i][j] = temp
		}
	}
}
