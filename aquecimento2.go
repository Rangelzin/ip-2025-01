package main

import "fmt"

func main() {
	var (
		peso float64 = 75.0
 		idade int = 18
		altura float64 = 1.76
		imc float64 
	)

	imc = peso / (altura * altura)

	fmt.Println("Olá")
	fmt.Println("Sua idade é igual a ", idade)
	fmt.Printf("Vc apresentou um peso de %.1f quilos e uma altura de %.2f metros\n",peso, altura)
	fmt.Printf("Vc apresentou um imc de %.2f\n" ,imc)
	
}