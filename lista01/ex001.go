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
		fmt.Printf("A Média do aluno é: %.2f\nAPROVADO!",media)
	} else {
		fmt.Printf("A Média do aluno é: %.2f\nREPROVADO!",media)
	}
}