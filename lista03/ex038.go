package main

import (
	f "fmt"
)

func main() {
	var (
		CPF                       string
		CPFnum                    = []int{}
		soma, j, digito1, digito2 int
	)

	f.Print("Digite o Número de CPF (somente números): ")
	f.Scan(&CPF)

	// Gurdar cada algarismo dentro de um slice
	for _, c := range CPF {
		digitoTOT := int(c - '0')
		CPFnum = append(CPFnum, digitoTOT)
	}

	for i := 10; i > 1; i-- {
		valorCPF := CPFnum[j]
		soma += (valorCPF * i)
		j++
	}
	resto := soma % 11

	if resto < 2 {
		digito1 = 0
	} else {
		digito1 = 11 - resto
	}

	j = 0
	soma = 0
	for i := 11; i > 1; i-- {
		valorCPF := CPFnum[j]
		soma += (valorCPF * i)
		j++
	}
	resto = soma % 11
	digito2 = 11 - resto

	if digito1 == CPFnum[9] && digito2 == CPFnum[10] {
		f.Printf("CPF: %s, válido!", CPF)
	} else {
		f.Printf("CPF: %s, inválido!", CPF)
	}
}
