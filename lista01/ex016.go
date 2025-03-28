package main

import (
	"fmt"
)

func main() {

	var (
		salary,newSalary float32
	)

	fmt.Print("Digite o valor o salário: R$")
	fmt.Scan(&salary)

	if salary <= 300 {
		newSalary =  salary * 1.5
	} else {
		newSalary =  salary * 1.3
	}

	fmt.Printf("Salário com reajuste R$%.2f\n", newSalary)
}