package main

import "fmt"


func main() {
	var numero int

	fmt.Print("Digite valor para verificação: ")
	fmt.Scan(&numero)

	if numero % 5 == 0 {
		fmt.Println("O valor é divisível por 5")
	} else {
		fmt.Println("O valor não é divisível por 5")
	}
}