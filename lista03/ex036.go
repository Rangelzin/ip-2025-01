package main

import f "fmt"

func main() {
	var (
		numero, resto int
		base16        = []int{}
	)

	f.Print("Digite um valor na base 10: ")
	f.Scan(&numero)

	for numero > 0 {
		resto = numero % 16
		base16 = append(base16, resto)
		numero = numero / 16
	}

	invertido := make([]int, len(base16))
	for i, j := 0, len(base16)-1; j >= 0; i, j = i+1, j-1 { // i++ e j--
		invertido[i] = base16[j]
	}

	f.Print("O número na base 16: ")
	for _, digito := range invertido {
		switch digito {
		case 10:
			f.Print("A")
		case 11:
			f.Print("B")
		case 12:
			f.Print("C")
		case 13:
			f.Print("D")
		case 14:
			f.Print("E")
		case 15:
			f.Print("F")
		default:
			f.Print(digito)
		}
	}
	f.Println()
}
