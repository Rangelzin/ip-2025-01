package main

import "fmt"

func main() {

	var (
		limiteCre,saldoInit, totDepo, retiradas float64
		conta 	int
	)

	fmt.Scan(&conta, &limiteCre, &saldoInit, &totDepo, &retiradas)

	Saldofinal := saldoInit + totDepo - retiradas
	
	if Saldofinal > limiteCre {
		fmt.Printf("Credito excedido: %.5f\n", Saldofinal)
	} else {
		fmt.Printf("Saldo: %.5f\n", Saldofinal)
	}
}