package main

import (
	f "fmt"
)

func main() {
	var numero, soma int

	f.Print("Digite ate que termo você deseja contar: ")
	f.Scan(&numero)

	for i := 1; i <= numero; i++ {
		soma += i
	}
	f.Printf("Soma: %d\n ", soma)
}
