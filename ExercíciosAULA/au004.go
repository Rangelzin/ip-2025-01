package main 

import "fmt"

func main() {

	fmt.Println(soma(10,20))
}

func soma(a,b int ) int {
	s := a + b
	return s
}