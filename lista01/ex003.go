package main 

import (
	"fmt"
	"strconv"
)

func main() {

	var n1,n2,n3 int

	fmt.Println("Insira 3 números: ")
	fmt.Scanln(&n1, &n2, &n3)

	if n1 >= 10 || n2 >= 10 || n3 >= 10 {
		fmt.Println("NUMERO INVÁLIDO!")
	} else {
		concatenado := strconv.Itoa(n1) + strconv.Itoa(n2) + strconv.Itoa(n3)
		numbigger, _ := strconv.Atoi(concatenado)
	
		
		fmt.Printf("%s,%d\n", concatenado, numbigger*numbigger)
	}
}