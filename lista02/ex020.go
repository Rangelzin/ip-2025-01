package main

import "fmt"

func main() {
	
	var (
		price,parcelas float32
		cod int
	)

	fmt.Print("Qual o preço da etiqueta e o código de pagamento: R$")
	fmt.Scan(&price, &cod)

	switch cod {
		case 1:
			price *= 0.9
			fmt.Println("Forma de pagamento: À vista, dinheiro ou cheque (10% de desconto)")
			fmt.Printf("Valor a ser pago = R$%.2f",price)
		case 2:
			price *= 0.95
			fmt.Println("Forma de pagamento: À vista, cartão de crédito (5% de desconto)")
			fmt.Printf("Valor a ser pago = R$%.2f",price)
		case 3:
			price = price
			parcelas = price/2
			fmt.Println("Forma de pagamento: Em 2 vezes,sem juros")
			fmt.Printf("Valor total a ser pago = R$%.2f, em 2x de R$%.2f",price,parcelas)
		case 4:
			price *= 1.10
			parcelas = price/3
			fmt.Println("Forma de pagamento: Em 3 vezes,com 10% de juros")
			fmt.Printf("Valor total a ser pago = R$%.2f, em 3x de R$%.2f",price,parcelas)
	}
}