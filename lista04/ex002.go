package main 

import f "fmt"

func main() {

	var (
		array10,array5 = []int{},[]int{}
		arrayR1,arrayR2 = []int{},[]int{}

		soma = 0
	)

	montagemVetor(&array10,10)
	montagemVetor(&array5,5)

	for _,v := range array10 {
		if v % 2 == 0 {
			soma = v
			for i := 0; i < 5; i++ {
				soma += array5[i]
			}
			arrayR1 = append(arrayR1, soma)
		} else {
			soma = v
			for i := 0; i < 5; i++ {
				soma += array5[i]
			}
			arrayR2 = append(arrayR2, soma)
		}
	}

	f.Printf("Primeiro vetor: %v\n",array10)
	f.Printf("Segundo vetor: %v\n",array5)

	f.Printf("Primeiro vetor resultante: %v\n",arrayR1)
	f.Printf("Segundo vetor resultante: %v\n",arrayR2)
}

func montagemVetor(lista *[]int, indice int) {

	f.Printf("VETOR COM %d\n",indice)
	for i := 0; i < indice; i++ {
		var num int
		f.Print("Digite um número: ")
		f.Scan(&num)
		*lista = append(*lista, num)
	}
}


