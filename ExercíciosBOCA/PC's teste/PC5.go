package main

import f "fmt"

func main() {
	var N,sum int

	f.Scan(&N)

	vet := make([]int, N)
	for i := 0; i < N; i++ {
		f.Scan(&vet[i])
		sum += vet[i]
	}

	media := sum / N

	for i := 0; i < N; i++ {
		if vet[i] > media {
			f.Printf("%d ", i)
		}
	}
}