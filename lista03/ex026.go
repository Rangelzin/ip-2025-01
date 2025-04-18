package main

import f "fmt"

func main() {
	var (
		soma      float64
		valorINIT = 100.0
	)

	for i := 0.0; i < 20; i++ {
		soma += (valorINIT - i) / fatorialFrac(int(i))
	}
	f.Printf("%.2f", soma)
}

func fatorialFrac(n int) float64 {
	resultado := 1.0
	for i := 2; i <= n; i++ {
		resultado *= float64(i)
	}
	return resultado
}
