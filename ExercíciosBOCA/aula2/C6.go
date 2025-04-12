package main

import "fmt"

func main() {

	var (
		reprovado, exame, aprovado, numAlu int
	)

	fmt.Scan(&numAlu)

	for i := 1; i <= numAlu; i++ {
		nota1, nota2 := 0.0, 0.0
		fmt.Scan(&nota1, &nota2)

		media := (nota1 + nota2) / 2

		if media <= 3 {
			fmt.Printf("Aluno %d: Reprovado\n", i)
			reprovado++
		} else if media > 3 && media < 7 {
			fmt.Printf("Aluno %d: Exame\n", i)
			exame++
		} else {
			fmt.Printf("Aluno %d: Aprovado\n", i)
			aprovado++
		}
	}
	fmt.Printf("Total Aprovados: %d\n", aprovado)
	fmt.Printf("Total Exame: %d\n", exame)
	fmt.Printf("Total Reprovados: %d", reprovado)
}
