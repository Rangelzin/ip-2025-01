package main

import "fmt"

func main() {
	
	var list = []int{12,221,32,454,412,4}
	x := 221

	fmt.Println(busca(list, x))
	
}

func busca( list []int, x int) int {
	for i := 0; i < len(list); i++ {
		if list[i] == x {
			return i
		}
	}

	return -1
}	