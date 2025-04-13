package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		pInit, D, deltaP, minP float64
		qtdInit, deltaQ        int
	)

	fmt.Scan(&pInit, &qtdInit, &D, &deltaP, &deltaQ, &minP)

	precoAtual := pInit
	qtdAtual := qtdInit

	// Larguras para cada coluna
	colPrecoWidth := 10
	colIngressosWidth := 20
	colLucroWidth := 10

	// Cabeçalho
	fmt.Printf("%-*s %s %-*s\n", colPrecoWidth, "Preco", centerText("Ingressos Vendidos", colIngressosWidth), colLucroWidth, "Lucro")

	for {
		if precoAtual < minP {
			break
		}

		Lucro := (precoAtual * float64(qtdAtual)) - D

		// Formata os valores e centraliza "Ingressos Vendidos"
		qtdText := fmt.Sprintf("%d", qtdAtual)
		fmt.Printf("%-*.*f %s %-*.*f\n",
			colPrecoWidth, 2, precoAtual,
			centerText(qtdText, colIngressosWidth),
			colLucroWidth, 2, Lucro)

		precoAtual -= deltaP
		qtdAtual += deltaQ
	}
}

// Função para centralizar texto
func centerText(text string, width int) string {
	padding := width - len(text)
	if padding <= 0 {
		return text
	}
	left := padding / 2
	right := padding - left
	return strings.Repeat(" ", left) + text + strings.Repeat(" ", right)
}
