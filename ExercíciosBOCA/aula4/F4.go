package main

import f "fmt"

func main() {
	var (
		N, M, X int
		test bool = true
	)
	f.Scan(&N, &M)

	// Cria a matriz 1 com o tamanho que coloca
	matriz := make([][]int, N)
	for i := 0; i < N; i++ {
		matriz[i] = make([]int, M)
		
	}
	// Armazena o valor da matriz1
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			f.Scan(&X)
			matriz[i][j] = X
		}
	}
	
	for i := 0; i < N; i++ {
		for j := 0; j < M-1; j++ {
			if matriz[i][j+1] < matriz[i][j] {
				test = false
			}
		}	
	}

	for j := 0; j < M; j++ {
		for i := 0; i < N-1; i++ {
			if matriz[i+1][j] < matriz[i][j] {
				test = false
			}
		}	
	}


	if test {
		f.Printf("SIM\n")
	} else {
		f.Printf("NAO\n")
	}
}


