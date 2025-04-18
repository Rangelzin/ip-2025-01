package main

import (
	f "fmt"
)

func main() {
	var base, expoente int

	f.Print("Digite o valor da base e do expoente: ")
	f.Scan(&base, &expoente)

	potencia := base
	for i := 1; i < expoente; i++ {
		if base >= 2 && expoente > 1 {
			potencia *= base
		} else {
			return
		}

	}
	f.Printf("O valor da potencia: %d", potencia)
}
