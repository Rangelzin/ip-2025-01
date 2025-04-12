package main

import "fmt"

func main() {
	var numero, soma int

	fmt.Scan(&numero)

	for i := 1; i <= numero; i++ {
		fmt.Printf("%d ", i)
		soma += i
	}
	fmt.Println()
	fmt.Print(soma)
}
