package main

import "fmt"

func main() {
	var (
		n int
		aprov, reprov, exame int		
		nota  []float64
		media []float64
		somaClass float64
	)
	
	fmt.Print("Quantos alunos? ")
	fmt.Scan(&n)
	
	nota = make([]float64, n*2)
	media = make([]float64, n) 

	for i := 0; i < n; i++ {
		fmt.Printf("Digite as notas do aluno %d: ", i+1)
		fmt.Scan(&nota[i*2], &nota[i*2+1])	

		media[i] = (nota[i*2] + nota[i*2+1]) / 2

		if media[i] >= 7 {
			aprov++
		} else if media[i] < 7 && media[i] >= 3 {
			exame++
		} else {
			reprov++
		}
		somaClass += media[i]
	}

	mediaClass := somaClass / float64(n)

	for i := 0; i < n; i++ {
		fmt.Printf("Aluno %d | Nota 1 = %.1f | Nota 2 = %.1f | Média = %.1f\n", i+1,nota[i*2], nota[i*2+1], media[i])
	}
	fmt.Printf("Total Aprovados: %d\n", aprov)
	fmt.Printf("Total Exame: %d\n", exame)
	fmt.Printf("Total Reprovados: %d\n", reprov)
	fmt.Printf("Media Classe: %.1f\n", mediaClass)
}