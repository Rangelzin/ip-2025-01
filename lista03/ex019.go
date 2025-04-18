package main

import f "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := i + 1; j <= 10; j++ {
			if j != 10 {
				f.Printf("(%d ,%d) ", i, j)
			} else {
				f.Printf("(%d ,%d)\n", i, j)
			}
		}
	}
}
