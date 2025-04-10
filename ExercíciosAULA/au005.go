package main

import "fmt"

func main() {
	
	fmt.Println(maior(1,2,3))
}

func maior(n1,n2,n3 int) int {
	
	var maior int

	for true {
		if n1 > maior {
			maior = n1
		} else if n2 > maior {
			maior = n2
		} else if n3 > maior {
			maior = n3
		} else {
			break
		}
	}

	return maior
}