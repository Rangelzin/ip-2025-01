package main

import (
	"fmt"
)

func main() {

	var nota float64

	fmt.Println("Digite o valor da nota: ")
	fmt.Scan(&nota)
	
	if nota < 6 {
		fmt.Printf("Nota: %.1f - Conceito: D\n", nota)
	} else if nota >=6 && nota < 7.5{
		fmt.Printf("Nota: %.1f - Conceito: C\n", nota)
	} else if nota >=7.5 && nota < 9 {
		fmt.Printf("Nota: %.1f - Conceito: B\n", nota)
	} else if nota >=9 && nota < 10{
		fmt.Printf("Nota: %.1f - Conceito: A\n", nota)
	} else {
		fmt.Print("Nota inválida!\n")
	}

}