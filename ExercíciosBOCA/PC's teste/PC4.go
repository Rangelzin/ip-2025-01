package main 

import "fmt"

func main() {
	var (
		N, sum float64
		sinal,valor float64 = 1.0,1.0
	)

	fmt.Scan(&N)

	for i := 1.0; i <= N; i++{
		sum += (1/ valor) * sinal
		sinal *= -1.0
		valor += 2
	} 
	
	numeroPI := sum * 4
	fmt.Printf("%.6f\n", numeroPI)
}