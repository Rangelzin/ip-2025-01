package main

import (
	f "fmt"
)

func main() {
	var N1, N2, quociente int

	f.Print("Digite dois numeros para calculo: ")
	f.Scan(&N1, &N2)

	resto := N1
	for i := 1; resto >= N2; i++ {
		resto -= N2
		quociente = i
	}
	f.Printf("Resto: %d | Quociente: %d", resto, quociente)
}
