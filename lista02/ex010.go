package main

import "fmt"

func main() {

	var destino,retorno int

	fmt.Print("Digite o número do seu destino: ")
	fmt.Scan(&destino)
	fmt.Print("Sua viagem vai ter retorno? ")
	fmt.Scan(&retorno)

	if destino < 1 || destino > 4 {
		fmt.Println("Destino inválido! Tente novamente.")
	} else if retorno < 1 || retorno > 2 {
		fmt.Println("Tipo de retorno inválido! Tente novamente.")
	} 

	if destino == 1 && retorno == 1 {
		fmt.Println("O seu destino é a Região Norte")
		fmt.Println("O preço da sua passagem é R$900,00")
	} else if destino == 1 && retorno == 2{
		fmt.Println("O seu destino é a Região Norte")
		fmt.Println("O preço da sua passagem é R$500,00")
	} else if destino == 2 && retorno == 1 {
		fmt.Println("O seu destino é a Região Nordeste")
		fmt.Println("O preço da sua passagem é R650,00")
	} else if destino == 2 && retorno == 2{
		fmt.Println("O seu destino é a Região Nordeste")
		fmt.Println("O preço da sua passagem é R$350,00")
	} else if destino == 3 && retorno == 1 {
		fmt.Println("O seu destino é a Região Centro-Oeste")
		fmt.Println("O preço da sua passagem é R$600,00")
	} else if destino == 3 && retorno == 2{
		fmt.Println("O seu destino é a Região Centro-Oeste")
		fmt.Println("O preço da sua passagem é R$350,00")
	} else if destino == 4 && retorno == 1 {
		fmt.Println("O seu destino é a Região Sul")
		fmt.Println("O preço da sua passagem é R$550,00")
	} else if destino == 4 && retorno == 2{
		fmt.Println("O seu destino é a Região Sul")
		fmt.Println("O preço da sua passagem é R$300,00")
	} 
}