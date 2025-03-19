package main 

import (
	"fmt"
	"strconv"
)

func main() {

	var n1,n2,n3 int
	var numero1,numero2,numero3 string

	fmt.Println("Insira 3 números: ")
	fmt.Scanln(&n1, &n2, &n3)

	if n1 >= 10 || n2 >= 10 || n3 >= 10 {
		fmt.Println("NUMERO INVÁLIDO!")
	} else {
		numero1 = strconv.Itoa(n1)
		numero2 = strconv.Itoa(n2)
		numero3 = strconv.Itoa(n3)
		
		fmt.Print(numero1,numero2, numero3)
		fmt.Println("")
	}
	
}