package main

import "fmt"

func main() {
	var a,b,c,d,det float32

	fmt.Println("Digite o valor dos quatro elementos da matriz: ")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)

	det = a*d - b*c
	fmt.Printf("O valor do determinante é = %.2f\n", det)
}