package main

import f "fmt"

func main() {
	soma := 0.0
	k := 1.0
	for i := 38.0; i >= 2; i-- {
		j := i - 1.0
		soma += (j * i) / k
		k++
	}
	f.Printf("%.2f", soma)
}
