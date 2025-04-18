package main

import (
	f "fmt"
	m "math"
)

func main() {
	numero := 6.3
	steps := m.Round(numero * 10)

	for i := 0.0; i <= steps; i++ {
		a := float64(i) / 10.0
		senoREAL := m.Sin(a)
		if a == numero {
			f.Printf("%.1f %.4f | seno real: %.4f", a, senoMaclaurin(a), senoREAL)
		} else {
			f.Printf("%.1f %.4f | seno real: %.4f\n", a, senoMaclaurin(a), senoREAL)
		}
	}
}
func fatorial(n int) float64 {
	resultado := 1.0
	for i := 2; i <= n; i++ {
		resultado *= float64(i)
	}
	return resultado
}

func senoMaclaurin(x float64) float64 {
	x = m.Mod(x, 2*m.Pi)
	if x > m.Pi {
		x = x - 2*m.Pi
	}
	sinal := 1.0
	if x < 0 {
		sinal = -1.0
		x = -x
	}
	if x > m.Pi/2 {
		x = m.Pi - x
	}
	return sinal * (x -
		powFatorial(x, 3) +
		powFatorial(x, 5) -
		powFatorial(x, 7))
}

func powFatorial(x, y float64) float64 {
	return m.Pow(x, y) / fatorial(int(y))
}
