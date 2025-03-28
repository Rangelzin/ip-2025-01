package main

import (
	"fmt"
	"math"
)

func main() {

	var a,h,areaB,vol float64

	fmt.Println("Digite o valor da altura e aresta (em metros): ")
	fmt.Scanln(&h, &a)

	areaB = (3 * a * a * math.Sqrt(3))/ 2
	vol = (areaB * h) / 3

	fmt.Printf("O volume da piramide é %.2f metros cúbicos\n", vol)
}