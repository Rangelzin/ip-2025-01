package main 

import "fmt"

func main() {

	var soma int

	for i := 1; i <= 100; i++ {
		fmt.Printf("%.d\n",i)
			soma += i
		}
		fmt.Printf("Soma é = %.d\n",soma)
}