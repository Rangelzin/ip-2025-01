package main

import f "fmt"

func main() {
	var a, b, N int

	f.Print("Digite os dois primeiros termos da sequência e quantidade de termos: ")
	f.Scan(&a, &b, &N)

	c := 0
	f.Printf("%d ", a)
	f.Printf("%d ", b)
	for i := 3; i <= N; i++ {
		if i%2 == 0 {
			c = b - a
		} else {
			c = b + a
		}
		f.Printf("%d ", c)
		a = b
		b = c
	}
}
