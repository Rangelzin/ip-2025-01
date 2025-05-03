package main

import f "fmt"

func main() {
	var numero int = 87
	binario := converterNum(numero)
	f.Println("O número na base 2:", binario)
}

func converterNum(x int) string {
	if x == 0 {
		return ""
	}
	return converterNum(x/2) + f.Sprintf("%d", x%2)
}
