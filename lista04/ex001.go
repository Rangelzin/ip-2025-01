package main 

import f "fmt"

func main() {
	var (
		array = []int{}
	)

	for i := 0; i < 10; i++ {
		var num int
		f.Print("Digite um número: ")
		f.Scan(&num)
		array = append(array, num)
	}
	
	for i := 0; i < 10; i++ {
		if array[i] > 50 {
			f.Printf("Posição: %d | Número: %d\n",array[i],i)
		}
	}
}