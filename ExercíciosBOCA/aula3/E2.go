package main

import (
	f "fmt"
	m "math"
)

func main() {
	var (
		qtd      int
		max, min = 0.0, m.MaxFloat64
	)

	f.Scan(&qtd)
	valores := make([]float64, qtd)
	for i := 0; i < qtd; i++ {
		f.Scan(&valores[i])
	}

	for _, v := range valores {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	deno, valoresNorm := max-min, 0.0
	if deno == 0.0 {
		for i := 0; i < qtd; i++ {
			f.Printf("%.2f ", deno)
		}
	} else {
		for _, v := range valores {
			valoresNorm = (v - min) / deno
			f.Printf("%.2f ", valoresNorm)
		}
	}

}
