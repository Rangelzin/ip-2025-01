package main

import f "fmt"

func main() {
	var A [10]int
	contagem := make(map[int]int)

	f.Println("Digite 10 números para incluir no vetor:")
	for i := 0; i < 10; i++ {
		f.Scan(&A[i])
		contagem[A[i]]++
	}

	f.Println("\nElementos repetidos:")
	for numero, freq := range contagem {
		if freq > 1 {
			f.Printf("%d aparece %d vezes\n", numero, freq)
		}
	}
}
