package main

import "fmt"

func main() {

	//var aprovados map[int]string
	//maps devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[454545453] = "Ana"
	aprovados[342432434] = "João"
	aprovados[545465656] = "Kiko"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[123456789])

	delete(aprovados, 123456789)

	fmt.Println(aprovados[123456789])

}
