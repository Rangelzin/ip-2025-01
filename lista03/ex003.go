package main

import "fmt"

func main() {
	var salarioCarlos float64

	fmt.Print("Digite o salario do Carlos: ")
	fmt.Scan(&salarioCarlos)
	
	salarioJoao := salarioCarlos / 3

	for i := 0; i < count; i++ {
		M := (salarioJoao * (1 + 0.02)) * i
	} 
}