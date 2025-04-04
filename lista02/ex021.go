package main

import "fmt"

func main() {

	var (
		id int
		nota1, nota2, nota3, mediaEx float64
		conceito,resultado string
	)

	fmt.Print("NÚMERO DE IDENTIFICAÇÃO: ")
	fmt.Scan(&id)
	fmt.Print("Digite o valor das 3 notas e a media dos exercicios: ")
	fmt.Scan(&nota1,&nota2,&nota3, &mediaEx)

	mediaFin := (nota1 + (nota2 * 2) + (nota3 * 3) + mediaEx) / 7

	if mediaFin > 0 && mediaFin < 4 {
		conceito = "E"
	} else if mediaFin >=4 && mediaFin < 6{
		conceito = "D"
	} else if mediaFin >=6 && mediaFin < 7.5{
		conceito = "C"
	} else if mediaFin >=7.5 && mediaFin < 9 {
		conceito = "B"
	} else if mediaFin >=9 && mediaFin < 10{
		conceito = "A"
	} else {
		fmt.Print("Média Final inválida!\n")
	}

	if conceito == "E" || conceito == "D" {
		resultado = "REPROVADO!"
	} else {
		resultado = "APROVADO!"
	}

	fmt.Printf("Numero do aluno = %d\n", id)
	fmt.Printf("Suas notas = %.1f, %.1f, %.1f\n", nota1,nota2,nota3)
	fmt.Printf("Média dos Exercícios = %.1f\n", mediaEx)
	fmt.Printf("Média Final = %.1f\n", mediaFin)
	fmt.Printf("Conceito = %s\n", conceito)
	fmt.Printf("Resultado = %s\n", resultado)
	
}