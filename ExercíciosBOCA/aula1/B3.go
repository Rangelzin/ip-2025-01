package main

import (
	"fmt"
)

func main() {

	var X,Y,Z int

	fmt.Scan(&X,&Y,&Z)

	if X + Y <= Z || X + Z <= Y || Y + Z <= X {
		fmt.Println("Nao forma triangulo")
	}else {
		if X == Y && Y == Z {
			fmt.Println("Equilatero")
		} else if X != Y && Y != Z{
			fmt.Println("Escaleno")
		} 
	}
}