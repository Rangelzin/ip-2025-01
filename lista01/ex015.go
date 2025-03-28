package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		num int
		val float64
	)

	fmt.Print("Digite o número para os quadrados pares: ")
	fmt.Scan(&num)

	if num > 5 && num < 2000 {
		for i := 2; i <= num; i += 2 {
		
			val = math.Pow(float64(i), 2)
	
			fmt.Printf("%d^2 = %.0f\n", i, val)
		}
	} else {
		fmt.Println("Número fora do intervalo, permitido apenas entre 5 a 2000.")
	}

	
}