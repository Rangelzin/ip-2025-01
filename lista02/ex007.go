package main

import "fmt"


func main() {
	var (
		number = [3]int{}
		maior,menor, inter int
	)

	for i := 0; i < 3; i++ {
		fmt.Printf("Digite o %d° valor: ",i+1)
		fmt.Scan(&number[i])
	}

	maior = number[0]
	menor = number[0]

	for i := 0; i < 3; i++ {
			
		if number[i] > maior { maior = number[i] }
		if number[i] < menor { menor = number[i] }

	}

	for i := 0; i < 3; i++ {
		if number[i] != maior && number[i] != menor {inter = number[i]}
	}

	fmt.Println(menor, inter, maior)
}