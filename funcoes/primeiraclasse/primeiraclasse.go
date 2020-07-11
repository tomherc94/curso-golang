package main

import "fmt"

var soma = func(a, b int) int { //funca global
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	sub := func(a, b int) int { //funcao local
		return a - b
	}

	fmt.Println(sub(2, 3))
}
