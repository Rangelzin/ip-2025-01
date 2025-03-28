package main

import "fmt"

func main() {

	var num int

	fmt.Print("Digite o numero que deseja verificar: ")
	fmt.Scan(&num)

	if  num % 5 == 0 && num % 3 == 0 {
		fmt.Println("O numero é divisivel")
	} else {
		fmt.Println("O numero não é divisivel")
	}
}
