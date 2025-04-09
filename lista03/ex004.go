package main 

import (
	"math"
	"fmt"
)

func main() {
	var numero,valorRaiz float64

	for {
		fmt.Print("Digite um valor: ")
		fmt.Scan(&numero)
	
		if numero > 0 {
			valorRaiz = math.Pow(numero, 0.5)
	
			if valorRaiz == math.Floor(valorRaiz) {
				fmt.Println("O número é um quadrado perfeito")
			} else {
				fmt.Println("O número NÃO é um quadrado perfeito")
			}
			
		} else {
			return 
		}
	}
}