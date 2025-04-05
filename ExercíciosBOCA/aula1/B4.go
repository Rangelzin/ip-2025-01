package main

import "fmt"

func main() {

	var idade int

	fmt.Scan(&idade)

	if idade >= 5 && idade <= 7 {
		fmt.Println("Infantil A")
	}else if idade >= 8 && idade <= 10 {
		fmt.Println("Infantil B")
	}else if idade >= 11 && idade <= 13 {
		fmt.Println("Juvenil A")
	}else if idade >= 14 && idade <= 17 {
		fmt.Println("Juvenil B")
	}else if idade >= 18  {
		fmt.Println("Adulto")
	}
}