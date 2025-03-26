package main

import "fmt"

func main() {

	var num int

	fmt.Println("Digite o numero que deseja verificar: ")
	fmt.Scan(&num)

	if  num % 5 == 0 && num % 3 == 0 {
		fmt.Println("O numero", num, "é divisivel")
	} else {
		fmt.Println("O numero ", num, " não é divisivel")
	}
}
