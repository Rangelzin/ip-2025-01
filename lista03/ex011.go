package main 

import (
	f"fmt"
	m"math"
)

func main() {
	var n float64

	for{
		f.Print("Digite um número para o fatorial: ")
		f.Scan(&n)
		if n == m.Floor(n) && n >= 0 {
			resultado := 1.0
			for i := n; i >= 2; i-- {
				resultado *= i
			}
			f.Printf("O fatorial de é: %.0f\n", resultado)
			break
		} else {
			f.Println("Número inválido, digite um número inteiro.")
			return 
		}
	}
}