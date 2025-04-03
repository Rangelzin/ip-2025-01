package main

import (
	"fmt"
)

func main() {

	var numero int

	fmt.Print("Digite um valor com 3 casas: ")
	fmt.Scan(&numero)

	if numero < 100 {
		fmt.Println("Seu número não tem 3 casas!")
	} else if numero > 999 {
		fmt.Println("Seu número tem mais de 3 casas!") 
	}else {
		centena := numero / 100
		dezena := (numero - centena*100) / 10
	
		fmt.Printf("Algarismo da dezena: %d\n", dezena)
	}
}