package main

import "fmt"

func main() {

	var (
		notas[5] float64
		soma   float64
	)
	for i := 0; i < len(notas); i++ {
		fmt.Printf("Digite a nota %d: ", i+1)
		fmt.Scan(&notas[i])

		soma += notas[i]
	}	
	fmt.Printf("A soma é: %.2f\n", soma)
}