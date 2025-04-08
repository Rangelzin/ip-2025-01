package main 

import "fmt"

func main() {
	var (
		numero int

	)
	fmt.Print("Digite um valor: ")
	fmt.Scan(&numero)

	if numero > 0 {
		for numero > 0 {
			fmt.Print("Digite um valor: ")
			fmt.Scan(&numero)

			quadrado := float64(numero * numero)

			for i := 1; i <= int(quadrado); i++ {
				if quadrado / float64(i) == float64(numero) {
					fmt.Println("Quadrado perfeito encontrado")
					break 
				}
			}

			fmt.Println("Loop encerrado para o número:", numero)
		}
	} else {
		fmt.Print("")
	}
}