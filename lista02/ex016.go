package main

import (
	"fmt"
	"math"
)

func main() {

	var A,B,C float32

	fmt.Print("Digite os coeficiente A, B e C da equação: ")
	fmt.Scan(&A, &B, &C)

	deter := (B * B) - (4*A*C)
	x1 := - B + float32(math.Sqrt(float64(deter))) / (2 * A)
	x2 := - B - float32(math.Sqrt(float64(deter)))/ (2 * A)

	if deter < 0 {
		fmt.Print("Raiz Imaginária/ Não real!")
	} else if deter == 0{
		fmt.Printf("Raiz Única! = %.0f", x1)
	} else {
		fmt.Printf("Raiz Dupla! = %.0f e %.0f", x1,x2)
	}
}