package main

import (
	"fmt"
)

func main() {
	var (
		altura, peso, sumAlt                         float64
		idade, resp, qtdPeopleElderly                int
		qtdPeopleHigh, qtdPeopleThin, contadorPeople int
	)
	for {

		fmt.Print("Digite a idade da pessoa: ")
		fmt.Scan(&idade)
		fmt.Print("Digite a altura da pessoa (em cm): ")
		fmt.Scan(&altura)
		fmt.Print("Digite a peso da pessoa (em KG): ")
		fmt.Scan(&peso)

		if idade > 50 {
			qtdPeopleElderly++
		}

		if idade > 10 && idade < 20 {
			sumAlt += altura
			qtdPeopleHigh++
		}

		if peso < 40 {
			qtdPeopleThin++
		}

		fmt.Print("Deseja continuar digitando dados? (1 - Sim, 2 - Não)? ")
		fmt.Scan(&resp)

		if resp == 2 {
			contadorPeople++
			mediaAlt := sumAlt / (float64(qtdPeopleHigh))
			percePeopleThin := (float64(qtdPeopleThin) / float64(contadorPeople) * 100)
			fmt.Printf("São %d pessoas com mais de 50 anos\n", qtdPeopleElderly)
			fmt.Printf("A média de altura das pessoas com mais de entre 10 e 20 anos é %.2f cm\n", mediaAlt)
			fmt.Printf("Porcentagem de pessoas com peso inferior a 40 quilos entre todas as pessoas analisadas é %.2f%%\n", percePeopleThin)
			return
		} else if resp == 1 {
			contadorPeople++
		}
	}
}
