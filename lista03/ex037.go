package main

import (
	f "fmt"
	s "strconv"
)

func main() {
	var numero int

	f.Print("Digite um valor na base 8: ")
	f.Scan(&numero)

	base8 := s.Itoa(numero)
	result := 0

	for _, c := range base8 {
		if c < '0' || c > '7' {
			f.Printf("Dígito inválido '%c' para base 8\n", c)
			return
		}
		result = result*8 + int(c-'0')
	}
	f.Printf("O número na base 10: %d\n", result)
}
