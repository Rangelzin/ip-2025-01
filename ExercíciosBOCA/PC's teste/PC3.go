package main 

import f"fmt"

func main() {
	N := 1

	f.Scan(&N)

	if N < 1 && N > 20{
		return 
	}	

	for i := 1; i <= N; i++ {
		for j := i; j <= N*i; j += i {
			f.Printf("%d ",j)
		}	
		f.Println()
	} 
}