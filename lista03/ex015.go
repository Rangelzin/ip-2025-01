package main

import f "fmt"

func main() {
	var Ntermos int

	f.Print("Digite a quantidade de termos: ")
	f.Scan(&Ntermos)

	for i := 1; i <= Ntermos; i++ {
		f.Printf("%d ", i*i)
	}
}
