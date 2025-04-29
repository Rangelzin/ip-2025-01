package main

import f "fmt"

func main() {

	var (
		alturas      = []float64{}
		altura, soma float64
		qtd          int
	)

	f.Scan(&qtd)
	for i := 0; i < qtd; i++ {
		f.Scan(&altura)
		alturas = append(alturas, altura)
		soma += altura
	}
	media := soma / float64(qtd)
	for _, alturai := range alturas {
		if alturai > media {
			f.Printf("%.2f\n", alturai)
		}
	}
}
