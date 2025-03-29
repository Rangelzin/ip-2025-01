package main 

import "fmt"

func main() {

	var numero int

	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	if numero > 20 && numero < 90 {
		fmt.Print("O seu número está no intervalo\n")
	} else {
		fmt.Print("O seu número não está no intervalo\n")
	}
}