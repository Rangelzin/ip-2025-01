package main

import "fmt"

func main() {
	var numero int

	for {
		fmt.Print("Digite um número positivo (ou 0 para sair): ")
		fmt.Scan(&numero)

		if numero <= 0 {
			break
		}
		encontrado := false

		for i := 1; i*i*i <= numero; i++ {
			if i*(i+1)*(i+2) == numero {
				encontrado = true
				break
			}
		}

		if encontrado {
			fmt.Println("O número é um produto de três números consecutivos (triangular especial)")
		} else {
			fmt.Println("O número NÃO é um produto de três números consecutivos")
		}
	}
}
