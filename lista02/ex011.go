package main

import "fmt"

func main() {

	var x,funcao float64

	fmt.Print("Digite o valor um valor de x (Sendo x diferente de 2): ")
	fmt.Scan(&x)

	if x == 2 {
		fmt.Println("F(x) indeterminada! Tente novamente.")
	} else {
		funcao = 8 / (2-x)
		fmt.Printf("Valor da F(x) = %.0f", funcao)
	}
}