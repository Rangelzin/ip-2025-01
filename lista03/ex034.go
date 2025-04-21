package main

import (
	f "fmt"
)

func main() {
	var N1, N2, resto int

	f.Print("Digite dois numeros para calculo: ")
	f.Scan(&N1, &N2)

	baseN1 := N1
	baseN2 := N2

	for {
		resto = N1 % N2
		N1 = N2
		N2 = resto
		if resto == 0 {
			break
		}
	}
	MDC := N1
	MMC := (baseN1 * baseN2) / MDC
	f.Printf("MMC: %d", MMC)
}
