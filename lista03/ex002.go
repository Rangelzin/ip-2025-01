package main

import "fmt"

func main() {

	var (
		soma  int
		media float64
	)

	for i := 52; i < 70; i = i + 2 {
		soma += i
	}
	media = float64(soma) / 10
	fmt.Println("Soma:", soma)
	fmt.Println("Média:", media)
}