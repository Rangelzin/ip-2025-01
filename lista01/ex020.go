package main

import (
	"fmt"
)

func main() {

	var (
		hr, min,totalseg int
		seg = [3]int{}
	)

	fmt.Println("Digite o valor de horas / minutos / segundos: ")
	fmt.Scan(&hr)
	fmt.Scan(&min)
	fmt.Scan(&seg[0])

	seg[1] = hr * 3600
	seg[2] = min * 60

	for i := 0; i < 3; i++ {
		totalseg += seg[i]
	}

	fmt.Printf("O tempo em segundos: %d\n", totalseg)
}