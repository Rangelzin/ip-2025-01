package main 

import f "fmt"

func main() {
	vetor := [5]int{2,3,4,5,6}
	resultado := recursao(vetor[:],len(vetor))
	f.Println(resultado)
}

func recursao(x []int,n int) int{
	if n < 1 {
		return 0
	}
	return x[n-1] + recursao(x,n-1)
}
