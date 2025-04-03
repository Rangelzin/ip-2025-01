package main

import "fmt"

func main() {
	var number1, number2, soma int

	fmt.Println("Digite dois valores para teste: ")
	fmt.Scan(&number1, &number2)

	soma = number1 + number2

	if soma > 20 {
		soma += 8
		fmt.Printf("Valor da soma ajustado = %d\n", soma)
	} else {
		soma -= 5
		fmt.Printf("Valor da soma ajustado = %d\n", soma)
	}
}