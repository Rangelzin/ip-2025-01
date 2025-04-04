package main

import (
	"fmt"
	"math"
)

func main() {

	var tipo int
	var volume, area,r,h float64


	fmt.Println("Qual tipo você deseja (1-cone / 2-cilindro / 3-esfera)?")
	fmt.Scan(&tipo)

	if tipo == 1 || tipo == 2 {
		fmt.Print("Digite o valor do raio e altura: ")
		fmt.Scan(&r, &h)
	} else {
		fmt.Print("Digite o valor do raio: ")
		fmt.Scan(&r)
	}

	switch tipo {
		case 1:
			volume = (math.Pi * r * r * h) / 3
			area = math.Pi * r * math.Sqrt(r*r + h*h)
		case 2:
			volume = math.Pi * r * r * h
			area = 2 * math.Pi * r * h
		case 3:
			volume = (4 * math.Pi * r * r * r) / 3
			area = 4 * math.Pi * r * r
	}

	fmt.Printf("Volume = %.2f metros cúbicos\n", volume)
	fmt.Printf("Área = %.2f metros quadrados\n", area)
}