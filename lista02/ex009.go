package main

import "fmt"

func main() {
	
	var buy, sale float64

	fmt.Print("Digite o valor de compra: R$")
	fmt.Scan(&buy)

	if buy < 10 {
       sale = buy * 1.7	
	   fmt.Printf("O Valor de venda com Lucros é de R$%.2f\n", sale)	
	} else if buy >= 10 && buy < 30{
       sale = buy * 1.5	
	   fmt.Printf("O Valor de venda com Lucros é de R$%.2f\n", sale)	
	} else if buy >= 30 && buy < 50{
       sale = buy * 1.4	
	   fmt.Printf("O Valor de venda com Lucros é de R$%.2f\n", sale)	
	} else {
       sale = buy * 1.3	
	   fmt.Printf("O Valor de venda com Lucros é de R$%.2f\n", sale)	
	}
}