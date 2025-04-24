package main

import "fmt"

func main() {
	nota := 0.0

	fmt.Scan(&nota)

	if nota >= 9 {
		fmt.Println("A")
	} else if nota >= 7 && nota < 9{
		fmt.Println("B")
	} else if nota >= 5 && nota < 7{
		fmt.Println("C")
	} else if nota >= 3 && nota < 5{
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}	