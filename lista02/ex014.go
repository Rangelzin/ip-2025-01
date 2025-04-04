package main

import "fmt"

func main() {

	var (
		valor float64
		respAR,respPIN,respVI,respDIR int
	)

	fmt.Print("Digite o valor do carro: R$")
	fmt.Scan(&valor)
	fmt.Print("O Carro possui Ar condicionado (R$ 1750,00)? ")
	fmt.Scan(&respAR)

	if respAR == 1 {valor += 1750}

	fmt.Print("O Carro possui Pintura metálica (R$ 800,00)? ")
	fmt.Scan(&respPIN)

	if respPIN == 1 {valor += 800}

	fmt.Print("O Carro possui Vidro elétrico (R$ 1200,00)? ")
	fmt.Scan(&respVI)

	if respVI == 1 {valor += 1200}

	fmt.Print("O Carro possui Direção hidráulica (c 2000,00)? ")
	fmt.Scan(&respDIR)

	if respDIR == 1 {valor += 2000}

	fmt.Printf("O valor final do carro é = R$%.2f", valor)

}	
