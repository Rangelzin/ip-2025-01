package main 

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um valor: ")
	fmt.Scan(&numero)

	if numero > 0 {
		for numero > 0 {
			fmt.Print("Digite um valor: ")
			fmt.Scan(&numero)
		}
	} else {
		fmt.Print("")
	}
}