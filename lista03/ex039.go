package main

import (
	f "fmt"
	m "math"
)

func main() {
	var (
		peso, maiorPeso       float64
		id, idPesado, idMagro int
		menorPeso             = m.MaxFloat64
	)

	for i := 1; i <= 90; i++ {
		f.Print("Digite o ID e o Peso do Boi em KG: ")
		f.Scan(&id, &peso)

		if peso > maiorPeso {
			maiorPeso = peso
			idPesado = id
		}
		if peso < menorPeso {
			menorPeso = peso
			idMagro = id
		}
	}

	f.Printf("Mais Gordo - ID: %d | Peso: %.1f\n", idPesado, maiorPeso)
	f.Printf("Mais Magro - ID: %d | Peso: %.1f\n", idMagro, menorPeso)
}
