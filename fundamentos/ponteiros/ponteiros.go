package main

import "fmt"

func main() {

	i := 1

	var p *int = nil

	p = &i //atribui o endereço de memória da variável i

	*p++ //desreferenciando (pegando o valor)

	i++

	//Go nao tem aritmética de ponteiros
	//p++

	fmt.Println(p, *p, i, &i)

}
