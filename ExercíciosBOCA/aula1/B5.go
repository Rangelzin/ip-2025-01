package main

import "fmt"

func main() {

	var (
		nota1, nota2, nota3 float64
		faltas 	int
	)

	fmt.Scan(&nota1, &nota2, &nota3, &faltas)

	media := (nota1 + nota2 + nota3) / 3

	if faltas * 4 > 64 {
		fmt.Println("Reprovado por falta")
	} else {
		if media >= 7 {
			fmt.Println("Aprovado")
		} else if media < 7 && media >= 5 {
			fmt.Println("Prova final")
		} else {
			fmt.Println("Reprovado")
		}
	}
}