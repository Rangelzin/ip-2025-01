package main

import "fmt"

var raio, altura,pi,area, custo float32

func main() {
	fmt.Println("Digite o valor do raio e a altura da lata(em metros):")
	fmt.Scan(&raio, &altura) 

	pi = 3.14159
	area = 2 * (pi * raio * raio) + (2 * pi * raio * altura)

	custo = area * 100

	fmt.Printf("O valor do custo é = R$%.2f\n", custo)
}