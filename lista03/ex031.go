package main

import f "fmt"

func main() {
	var contador, soma float64 = 1, 1

	for i := 1; i <= 64; i++ {
		f.Printf("%d - %.0f\n", i, contador)
		soma += contador
		contador *= 2
	}
	f.Printf("A soma de todas os milhos presentes no tabuleiro - %.0f\n", soma)
}
