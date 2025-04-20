package main

import (
	f "fmt"
	m "math"
)

func main() {

	for i := 0.0; i <= 20; i += 0.5 {
		volume := (4 * m.Pi * m.Pow(i, 3)) / 3
		f.Printf("%.1f - %.3f\n", i, volume)
	}
}
