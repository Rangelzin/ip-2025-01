package main 

import "fmt"

func main() {

	var numero int

	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	if numero > 0 {
		fmt.Print("O seu número é positivo\n")
	} else if numero < 0 {
		fmt.Print("O seu número é negativo\n")
	}else {
		fmt.Print("O seu número é nulo\n")
	}
}