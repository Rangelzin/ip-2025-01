package main 

import "fmt"

func main() {

	var (
		jogos int
		 pubtot float32
		 popular float32
		 geral float32
		 arquibancada float32
		 cadeiras float32
		 
	)
	receita := [4]float32{1,2,3,4}
	receitatot := [3]float32{1,2,3}

	fmt.Print("São quantos jogos? ")
	fmt.Scanln(&jogos)

	for i := 0; i <= jogos-1; i++ {
		
		fmt.Printf("O número de pessoas que compraram ingresso para o jogo N. %.d: ", i+1 )
		fmt.Scanln(&pubtot)
		fmt.Printf("A percentagem de pessoas que compraram ingresso na categoria Popular do jogo N. %.d: ", i+1 )
		fmt.Scanln(&popular)
		fmt.Printf("A percentagem de pessoas que compraram ingresso na categoria Geral do jogo N. %.d: ", i+1 )
		fmt.Scanln(&geral)
		fmt.Printf("A percentagem de pessoas que compraram ingresso na categoria Arquibancada do jogo N. %.d: ", i+1 )
		fmt.Scanln(&arquibancada)
		fmt.Printf("A percentagem de pessoas que compraram ingresso na categoria Cadeiras do jogo N. %.d: ", i+1 )
		fmt.Scanln(&cadeiras)
		
		receita[0] = pubtot * (popular/100) * 1
       	receita[1] = pubtot * (geral/100) * 5
       	receita[2] = pubtot * (arquibancada/100) * 10
       	receita[3] = pubtot * (cadeiras/100) * 20

		receitatot[i] = receita[0] + receita[1] + receita[2] + receita[3]
	}
	
	for i := 0; i < jogos; i++ {
		fmt.Printf("A RENDA DO JOGO N. %d É DE: R$ %.2f\n", i+1, receitatot[i])
	}
	
}