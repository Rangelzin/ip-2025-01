package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		pInit   = 6.0
		D       = 300.00
		deltaP  = 0.6
		minP    = 1.00
		qtdInit = 130.0
		deltaQ  = 30.0
	)

	precoAtual := pInit
	qtdAtual := qtdInit
	maiorLucro := 0.0
	faixaPreco := 0.0
	faixaIng := 0.0

	fmt.Printf("%-5s %s %-10s\n",
		"Preço",
		centralizarTexto("Ingressos Vendidos", 20),
		"Lucro",
	)

	for {
		Lucro := (precoAtual * qtdAtual) - D
		if precoAtual > minP || precoAtual == minP {
			qtdText := fmt.Sprintf("%.0f", qtdAtual)

			fmt.Printf("%-5.2f %s %-10.2f\n",
				precoAtual,
				centralizarTexto(qtdText, 20),
				Lucro,
			)

			if Lucro > maiorLucro {
				maiorLucro = Lucro
				faixaPreco = precoAtual
				faixaIng = qtdAtual
			}
			precoAtual -= deltaP
			qtdAtual += deltaQ
		} else {
			fmt.Printf("Lucro maximo: %.2f\n", maiorLucro)
			fmt.Printf("Na faixa de preco: %.2f com %.0f ingressos.", faixaPreco, faixaIng)
			return
		}
	}
}

func centralizarTexto(texto string, largura int) string {
	espacos := largura - len(texto)
	if espacos <= 0 {
		// Se o texto for maior que a largura, apenas retorna o texto original.
		return texto
	}
	// Calcula quantos espaços vão à esquerda e à direita.
	esquerda := espacos / 2
	direita := espacos - esquerda
	return strings.Repeat(" ", esquerda) + texto + strings.Repeat(" ", direita)
}
