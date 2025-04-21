package main

import f "fmt"

func main() {
	var (
		numero, resto int
		base2         = []int{}
	)

	f.Print("Digite um valor na base 10: ")
	f.Scan(&numero)

	for numero > 0 {
		resto = numero % 2
		base2 = append(base2, resto)
		numero = numero / 2
	}

	invertido := make([]int, len(base2))
	for i, j := 0, len(base2)-1; j >= 0; i, j = i+1, j-1 { // i++ e j--
		invertido[i] = base2[j]
	}

	f.Print("O número na base 2: ")
	for _, digito := range invertido {
		f.Print(digito)
	}
}
