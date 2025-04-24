package main 

import f "fmt"

func main() {
	var (
		array = []int{}
		somaPair,qtdOdd = 0,0

		arrayPair,arrayOdd = []int{},[]int{}
	)

	for i := 0; i < 10; i++ {
		var num int
		f.Print("Digite um número: ")
		f.Scan(&num)
		array = append(array, num)
	}

	for _,v := range array {
		if v % 2 == 0 {
			arrayPair = append(arrayPair, v)
			somaPair += v
		} else {
			arrayOdd = append(arrayOdd, v)
			qtdOdd++
		}
	}

	f.Printf("Número Pares digitados: %v\n",montagemVetor(arrayPair))
	f.Printf("Soma dos números pares: %v\n",somaPair)
	f.Printf("Número Impares digitados: %v\n",montagemVetor(arrayOdd))
	f.Printf("A quantidade de números ímpares digitados: %v\n",qtdOdd)
}

func montagemVetor(lista []int) string {
	resultado := ""
	for _, digito := range lista {
		resultado += f.Sprintf("%d ", digito)
	}
	return resultado
}