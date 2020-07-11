package main

import "fmt"

func inc1(n int) {
	n++
}

//revisao: um ponteiro é representado por um *
func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja, ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) //passagem por valor (passa uma cópia)
	fmt.Println(n)

	// revisão: & usado para obter o endereço da variável

	inc2(&n)
	fmt.Println(n)

}
