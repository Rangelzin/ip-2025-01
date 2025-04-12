package main

import (
	"fmt"
	"math"
)

func main() {
	var numero float64

	fmt.Scan(&numero)
	encontrado := false

	for i := 1.0; math.Pow(i, 3) <= numero; i++ {
		forms := i * (i + 1) * (i + 2)

		if forms == numero {
			encontrado = true
			break
		}

	}

	if encontrado {
		fmt.Printf("%.0f eh triangular", numero)
	} else {
		fmt.Printf("%.0f nao eh triangular", numero)
	}
}
