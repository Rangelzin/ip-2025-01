package main

import "fmt"

func main() {
	var (
		soma,n float64
		qtdA int
	)
	
	fmt.Println("Digite a quantidade de alunos: ")
	fmt.Scan(&qtdA)

	if qtdA <= 0 {
		fmt.Println("Quantidade inválida de alunos!")
	} else {
		for i := 0; i < qtdA; i++ {
			fmt.Printf("Qual a nota do aluno %d°: ",i+1)
			fmt.Scan(&n)
			soma += n
		}
	
		media := soma / float64(qtdA)
	
		fmt.Printf("A média global dos alunos é %.2f\n",media)
	}
}