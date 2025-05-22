package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	// Ler matriz
	mat := make([][]int, n)
	montarMatriz5(&mat, n, m)
	guardarValor5(&mat, n, m)

	// Limites da espiral
	top, bottom := 0, n-1
	left, right := 0, m-1

	for top <= bottom && left <= right {

		for j := left; j <= right; j++ {
			fmt.Print(mat[top][j], " ")
		}
		top++

		for i := top; i <= bottom; i++ {
			fmt.Print(mat[i][right], " ")
		}
		right--

		if top <= bottom {
			for j := right; j >= left; j-- {
				fmt.Print(mat[bottom][j], " ")
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				fmt.Print(mat[i][left], " ")
			}
			left++
		}
	}
}

func montarMatriz5(l *[][]int, X, Y int) {
	for i := 0; i < X; i++ {
		(*l)[i] = make([]int, Y)
	}
}

func guardarValor5(l *[][]int, X, Y int) {
	var temp int
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			fmt.Scan(&temp)
			(*l)[i][j] = temp
		}
	}
}
