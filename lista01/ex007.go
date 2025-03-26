package main

import "fmt"

func main() {
	var (
		tempF,tempC,pol,mm  float32
	)

	fmt.Println("Informe a temperatura em Fahrenheit e depois a quantidade de chuva em polegadas: ")
	fmt.Scan(&tempF)
	fmt.Scan(&pol)

	tempC = 5 * (tempF - 32)/9
	mm = pol * 25.4

	fmt.Printf("O valor em Celsius = %.2f°C\n", tempC)
	fmt.Printf("A quantidade de chuva e = %.2f mm\n", mm)
}