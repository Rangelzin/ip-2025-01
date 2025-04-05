package main

import "fmt"

func main() {
	var (
		matricula int
		horasExtras float64
	)
	fmt.Print("Digite a matrícula do funcionário: ")
	fmt.Scan(&matricula)
	fmt.Print("Foram quantas horas-extaras? ")
	fmt.Scan(&horasExtras)
	
	salarioMinimo := 788.00
	salarioExtra := horasExtras * 10.00
	salarioBruto := (3 * salarioMinimo) + salarioExtra

	if salarioBruto	> 1500 {salarioBruto *= 0.88}
	if salarioBruto > 2000 {salarioBruto *= 0.80}

	fmt.Printf("O salário líquido do funcionário %d é R$ %.2f\n", matricula, salarioBruto)
}