package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//forma serial
	//fale("Maria", "Pq vc nao fala comigo", 3)
	//fale("Joao", "Só posso faalr dps de vc", 1)

	//go fale("Maria", "Ei...", 500)
	//go fale("Joao", "Opa...", 500)
	//time.Sleep(time.Second * 5)
	//fmt.Println("fim")

	go fale("Maria", "Entendi", 10)
	fale("Joao", "Parabéns", 5)

}
