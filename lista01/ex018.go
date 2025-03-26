package main

import (
	"fmt"
)

func main() {

	var (
		valInit, r , n,soma int
	)

	fmt.Print("Digite o valor inicial / razão / qtd de elementos: ")
	fmt.Scan(&valInit, &r, &n)

	for i := 1; i <= n; i++ {
		soma += valInit + ((i-1)*r)
	}
	fmt.Printf("A soma da P.A é %d\n", soma)
}