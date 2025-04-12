package main

import (
	"fmt"
	"math"
)

func main() {
	var numero float64
	fmt.Scan(&numero)

	steps := math.Round(numero * 10)

	for i := 0.0; i <= steps; i++ {
		a := float64(i) / 10.0
		if a == numero {
			fmt.Printf("%.1f %.4f", a, senoMaclaurin(a))
		} else {
			fmt.Printf("%.1f %.4f\n", a, senoMaclaurin(a))
		}
	}
}
func fatorial(n int) float64 {
	resultado := 1.0
	for i := 2; i <= n; i++ {
		resultado *= float64(i)
	}
	return resultado
}

func senoMaclaurin(x float64) float64 {
	return x -
		powFatorial(x, 3) +
		powFatorial(x, 5) -
		powFatorial(x, 7)
}

func powFatorial(x, y float64) float64 {
	return math.Pow(x, y) / fatorial(int(y))
}
