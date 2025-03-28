package main

import "fmt"


func main() {

	var qtd int
	tempF :=  [3]float32{}
	tempC :=  [3]float32{}
	
	fmt.Print("Quantas temperaturas deseja informar? ")
	fmt.Scan(&qtd)

	for i := 0; i < qtd; i++ {
		fmt.Print("Qual a temperatura Nº", i+1, ": ")
		fmt.Scan(&tempF[i])
		
		tempC[i] = 5 * (tempF[i] - 32)/9
	}
	for i := 0; i < qtd; i++ {
		fmt.Printf("%.2f Fahrenheit equivale a %.2f Celsius\n", tempF[i], tempC[i])
	}
}