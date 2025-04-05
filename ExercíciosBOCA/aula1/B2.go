package main

import (
	"fmt"
)

func main() {

	var numero int

	fmt.Scan(&numero)

	if numero % 2 == 0 {
		fmt.Printf("%d Par\n", numero)
	} else {
		fmt.Printf("%d Impar\n", numero)
	}

}