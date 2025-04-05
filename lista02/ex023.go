package main

import "fmt"

func main() {
	var idade int

	fmt.Print("Informe a idade do cidadão: ")
	fmt.Scan(&idade)
	
	if idade < 16 {
		fmt.Println("-> Não-eleitor (abaixo de 16 anos)")
	}else if idade >= 16 && idade < 18 || idade > 65 {
		fmt.Println("-> Eleitor facultativo (entre 16 e 18 anos e maior de 65 anos).")
	}else if idade >= 18 {
		fmt.Println("-> Eleitor obrigatório (entre 18 e 65 anos)")
	}
}