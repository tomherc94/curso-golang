package main

import "fmt"

func main() {

	funcsESalarios := map[string]float64{
		"José João":      11233.45,
		"Gabriela Silva": 10567.90,
		"Pedro Junior":   1200.00,
	}

	funcsESalarios["Rafael Filho"] = 1350.00

	delete(funcsESalarios, "inexistente") //excluir item que nao existe

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
