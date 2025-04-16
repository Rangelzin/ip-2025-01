package main

import (
	"fmt"
)

func main() {

	l := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(BuscaBinaria(l, 9))
}

func BuscaBinaria(l []int, x int) int {
	n := len(l)
	e := 0
	d := n - 1
	for e <= d {
		m := (e + d) / 2
		if l[m] == x {
			return m
		} else if l[m] < x {
			e = m + 1
		} else {
			d = m - 1
		}
	}
	return -1
}
