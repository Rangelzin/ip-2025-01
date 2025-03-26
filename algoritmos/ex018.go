package main

import (
	"fmt"
)

func main() {

	var (
		valInit, r , n int
	)

	fmt.Print("Digite o valor inicial / razão / qtd de elementos: ")
	fmt.Scan(&valInit, &r, &n)

	valEnd := valInit + ((n-1)*r)

	soma := ((valInit + valEnd)*n) / 2 

	fmt.Printf("A soma da P.A é %d\n", soma)
}