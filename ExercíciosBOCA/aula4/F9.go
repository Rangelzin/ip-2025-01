package main

import "fmt"

func main() {
	// Declaração das variáveis para dimensões da matriz
	var N, M int

	// Lê as dimensões N (linhas) e M (colunas) da matriz
	fmt.Scan(&N, &M)

	// Inicializa a matriz A com N linhas
	A := make([][]int, N)
	montarMatriz4(&A, N, M)
	guardarValor4(&A, N, M)

	// Inicializa o filtro F com 3 linhas (matriz 3x3)
	F := make([][]int, 3)
	montarMatriz4(&F, 3, 3)
	guardarValor4(&F, 3, 3)

	// Processa a convolução válida
	// Percorre as linhas onde o filtro cabe completamente (de 1 até N-2)
	for i := 1; i <= N-2; i++ {
		// Percorre as colunas onde o filtro cabe completamente (de 1 até M-2)
		for j := 1; j <= M-2; j++ {
			sum := 0 // Inicializa o acumulador para esta posição

			// Aplica o filtro 3x3 na posição atual (i,j)
			// Percorre os deslocamentos verticais (di) de -1 a +1
			for di := -1; di <= 1; di++ {
				// Percorre os deslocamentos horizontais (dj) de -1 a +1
				for dj := -1; dj <= 1; dj++ {
					// Verifica se está dentro dos limites da matriz (redundante para convolução válida)
					if i+di >= 0 && i+di < N && j+dj >= 0 && j+dj < M {
						// Acumula o produto entre o elemento da matriz e o correspondente do filtro
						// di+1 e dj+1 convertem os índices do filtro de [-1,1] para [0,2]
						sum += A[i+di][j+dj] * F[di+1][dj+1]
					}
				}
			}

			// Formatação da saída:
			// Se for o primeiro elemento da linha, imprime sem espaço antes
			if j == 1 {
				fmt.Print(sum)
			} else {
				// Para outros elementos, imprime com espaço antes
				fmt.Printf(" %d", sum)
			}
		}
		// Após processar toda uma linha, imprime uma quebra de linha
		fmt.Println()
	}
}

func montarMatriz4(l *[][]int, X, Y int) {
	for i := 0; i < X; i++ {
		(*l)[i] = make([]int, Y)
	}
}

func guardarValor4(l *[][]int, X, Y int) {
	var temp int
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			fmt.Scan(&temp)
			(*l)[i][j] = temp
		}
	}
}
