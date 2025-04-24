package main

import "fmt"

func main() {
	ano := 0

	fmt.Scan(&ano)

	animal := (ano % 12) + 1

	switch animal {
		case 1:
			fmt.Println("Macaco")
		case 2:
			fmt.Println("Galo")
		case 3:
			fmt.Println("Cao")
		case 4:
			fmt.Println("Porco")
		case 5:
			fmt.Println("Rato")
		case 6:
			fmt.Println("Boi")
		case 7:
			fmt.Println("Tigre")
		case 8:
			fmt.Println("Coelho")
		case 9:
			fmt.Println("Dragao")
		case 10:
			fmt.Println("Serpente")
		case 11:
			fmt.Println("Cavalo")
		case 12:
			fmt.Println("Cabra")
	}
}	