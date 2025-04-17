package main

import f "fmt"

func main() {
	var N1, N2 int

	f.Print("Digite o valor de N1 e N2: ")
	f.Scan(&N1, &N2)

	for i := N1; i <= N2; i++ {
		primo := true
		if i < 2 {
			primo = false
		} else {
			for j := 2; j*j <= i; j++ {
				if i%j == 0 {
					primo = false
					break
				}
			}
		}

		if primo {
			f.Printf("O número %d é primo\n", i)
		} else {
			f.Printf("O número %d NÃO é primo\n", i)
		}
	}
}
