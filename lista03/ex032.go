package main

import f "fmt"

func main() {
	var N1, N2, multi float64

	f.Print("Digite dois valores que você queria multiplicar: ")
	f.Scan(&N1, &N2)

	for i := 1.0; i <= N2; i++ {
		multi += N1
	}
	f.Printf("%.0f x %.0f = %.0f", N1, N2, multi)
}
