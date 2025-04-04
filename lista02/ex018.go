package main

import "fmt"

func main() {

	var	dia,cat int
	var	price float64
	
	fmt.Print("Que dia da semana é hoje? ")
	fmt.Scan(&dia)
	fmt.Print("Qual o preço do DVD normal? R$")
	fmt.Scan(&price)
	fmt.Print("Qual a categoria do DVD? ")
	fmt.Scan(&cat)

	if dia == 2 || dia == 3 || dia == 5 {price *= 0.6}
	if cat == 2 {price *= 1.15}

	fmt.Printf("O Preço final do DVD é: R$%.2f\n",price)
}