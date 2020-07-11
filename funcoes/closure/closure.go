package main

import "fmt"

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()

	imprimeX() //imprimi a variavel do escopo de origem, ou seja, o valor 10
}
