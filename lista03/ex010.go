package main

import(
	f"fmt"
)

func main() {
	var n int

	f.Print("Digite a quantidade de valores da Fibonacci: ")
	f.Scan(&n)

	a, b := 0, 1
	
	for i := 0; i < n; i++ {
		if i == 0 {
			f.Printf("%d ", a)
		} else if i == 1 {
			f.Printf("%d ", b)
		} else {
			c := a + b
			f.Printf("%d ", c)
			a = b
			b = c
		}
	}
	f.Print("\n")
}
		