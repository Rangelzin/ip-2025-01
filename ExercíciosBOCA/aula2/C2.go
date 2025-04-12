package main

import (
	"fmt"
)

func main() {

	var (
		l, r, soma int
		contador   float64
	)

	fmt.Scan(&l, &r)

	if l == r {
		soma, contador = 0, 1
	} else {
		if l%2 == 0 {
			for i := l; i <= r; i = i + 2 {
				soma += i
				contador++

			}
		} else {
			for i := l + 1; i <= r; i = i + 2 {
				soma += i
				contador++
			}
		}
	}

	media := float64(soma) / contador
	fmt.Println(soma)
	fmt.Print(media)
}
