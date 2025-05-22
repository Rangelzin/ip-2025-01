package main

import f "fmt"

func main() {
	var (
		N,Q int
		valid bool = false
	)

	f.Scan(&N)

	vetN := make([]int, N)
	for i := 0; i < N; i++ {
		f.Scan(&vetN[i])
	}
	f.Scan(&Q)


	valQ := make([]int,Q)
	for i := 0; i < Q; i++ {
		f.Scan(&valQ[i])
	}

	for i := 0; i < Q; i++ {
		for j := 0; j < N; j++ {
			if valQ[i] == vetN[j]{
				valid = true
				break
			} else {
				valid = false 
			}
		}	
		if valid {
			f.Printf("SIM\n")
		} else {
			f.Printf("NAO\n")
		}
	}
}