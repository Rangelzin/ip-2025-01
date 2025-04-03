package main

import "fmt"


func main() {
	var number1, number2 int

	fmt.Println("Digite 2 valores para verificação: ")
	fmt.Scan(&number1, &number2)

	if number1 % number2 == 0 {
		fmt.Printf("O valor %d é divisível por %d\n", number1 ,number2)
	} else {
		fmt.Printf("O valor %d não é divisível por %d\n", number1 ,number2)
	}
}