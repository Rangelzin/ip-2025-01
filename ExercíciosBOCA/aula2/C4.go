package main

import (
	"fmt"
	"math"
)

func main() {

	var numero int
	for {
		fmt.Scan(&numero)
		if numero > 0 {
			valorRaiz := math.Pow(float64(numero), 0.5)

			if valorRaiz == math.Floor(valorRaiz) {
				fmt.Printf("%d eh quadrado perfeito\n", numero)
			} else {
				fmt.Printf("%d nao eh quadrado perfeito\n", numero)
			}
		} else {
			return
		}
	}
}
