package main

import "fmt"

func main() {
	var base float64
	var expoente int
	
	fmt.Print("Digite o valor da base e do expoente: ")
	fmt.Scan(&base, &expoente)

	resultado := 1.0
	for i := 0; i < expoente; i++ {
		resultado *= base
	}
	fmt.Printf("O resultado de %.0f elevado a %d é %.0f\n", base, expoente, resultado)
}
