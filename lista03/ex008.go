package main 

import "fmt"

func main() {
	var soma int

	for i := 1; i <= 20; i++ {
		soma += i
		fmt.Printf("%d\n",i)
	}
	fmt.Printf("Soma: %d\n",soma)
}