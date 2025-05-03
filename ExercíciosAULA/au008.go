package main 

import f "fmt"

func main() {
	var fatorial int = 12
	resultado := recursao(fatorial)
	f.Println(resultado)
}

func recursao(x int) int{
	if x == 0 {
		return 1
	}
	return x * recursao(x-1)
}
