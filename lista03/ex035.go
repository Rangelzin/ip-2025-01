package main

import f"fmt"

func main() {
	var numero int
	
	f.Print("Digite um valor na base 10: ")
	f.Scan(&numero)

	binario := ""
	for numero > 0 {
		binario = f.Sprintf("%d", numero%2) + binario
		numero = numero / 2
	}

    f.Println("O número na base 2:", binario)
}
