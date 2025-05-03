package main 

import f "fmt"

func main() {
	base, expoente := 16,4
	resultado := recursao(base,expoente)
	f.Println(resultado)
}

func recursao(x int,n int) int{
	if n < 1 {
		return 1
	}
	return x * recursao(x,n-1)
}
