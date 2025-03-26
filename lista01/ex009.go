package main

import "fmt"

func main() {
	var a, b, c, delta float32

	fmt.Println("Digite os coeficientes A,B e C: ")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)


	delta = b*b - 4*a*c

	fmt.Printf("O valor do delta é = %.2f\n", delta)	
}