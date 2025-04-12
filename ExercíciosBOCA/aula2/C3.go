package main

import (
	"fmt"
	"math"
)

func main() {

	var base, expoente float64

	fmt.Scan(&base, &expoente)

	potencia := math.Pow(base, expoente)

	fmt.Printf("%.0f", potencia)
}
