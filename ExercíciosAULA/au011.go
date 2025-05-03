package main 

import (
	f "fmt"
)
func main() {
	vetor := [5]int{321,1321,3212,423,534}
	recursao(vetor[:],len(vetor))
	f.Printf("%v\n",vetor)
}

func recursao(x []int, n int) {
    if n <= 1 {
        return 
    }
    x[0], x[n-1] = x[n-1], x[0]
	recursao(x[1:n-1], n-2)
}



