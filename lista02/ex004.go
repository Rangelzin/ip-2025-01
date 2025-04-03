package main

import (
	"fmt"
	"math"
)

func main() {
	var numero float64

	fmt.Print("Digite um valor qualquer: ")
	fmt.Scan(&numero)

	if numero >= 0 {
		numero = math.Sqrt(numero)
		fmt.Println(numero)
	} else {	
		numero *= numero
		fmt.Println(numero)
	}
}