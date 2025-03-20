package main 

import (
	"fmt"
)

func main() {

	var media float32
	notas := [3]float32{1,2,3}
	
	fmt.Println("Envie as 3 notas: ")
	fmt.Scanln(&notas[0],&notas[1],&notas[2])

	media = (notas[0]+ notas[1]+ notas[2]) / 3.0

	if media >= 6 {
		fmt.Println(fmt.Sprintf("A Média do aluno é: "),media)
		fmt.Println("APROVADO!")
	} else {
		fmt.Println(fmt.Sprintf("A Média do aluno é: "),media)
		fmt.Println("REPROVADO!")
	}
}