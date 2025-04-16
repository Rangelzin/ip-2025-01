package main

import "fmt"

func main() {
	var (
		numero, soma, media, maior, menor, somaPar, somaImpar, perceImp, mediaPar float64
	)
	contador := 1.0
	i := 0.0
	for {
		fmt.Print("Digite um valor: ")
		fmt.Scan(&numero)
		if numero != 30000 {
			soma += numero
			media = soma / contador

			if int(numero)%2 == 0 {
				somaPar += numero
				i++
			} else {
				somaImpar += numero
			}
			if numero > maior {
				maior = numero
			}
			if numero < menor {
				menor = numero
			}
			perceImp = (somaImpar / soma) * 100
			mediaPar = somaPar / i
			contador++
		} else {
			fmt.Printf("A soma dos números digitados: %.2f\n", soma)
			fmt.Printf("A quantidade de números digitado: %.2f\n", contador)
			fmt.Printf("A média dos números digitados: %.2f\n", media)
			fmt.Printf("O maior número digitado: %.0f\n", maior)
			fmt.Printf("O menor número digitado: %.0f\n", menor)
			fmt.Printf("A média dos números pares: %.2f\n", mediaPar)
			fmt.Printf("A percentagem dos números ímpares entre todos os números digitados: %.2f%%\n", perceImp)
			return
		}
	}
}
