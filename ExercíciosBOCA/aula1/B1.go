package main

import (
	"fmt"
	"math"
)

func main() {

	var A,B,C float64

	fmt.Scan(&A,&B,&C)

	deter := (B * B) - (4*A*C)
	x1 := (-B + math.Sqrt(deter)) / (2 * A)
	x2 := (-B - math.Sqrt(deter)) / (2 * A)

	if deter < 0 {
		fmt.Println("Nao ha raizes reais")
	} else if A == 0{
		fmt.Println("Nao e equacao do segundo grau")
	} else {
		fmt.Printf("%.5f %.5f\n",x1,x2)
	}
}