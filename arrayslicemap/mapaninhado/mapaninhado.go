package main

import "fmt"

func main() {

	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva":  12000.00,
			"Guilherme Jonas": 1000.00,
		},
		"J": {
			"João Valerio": 1300.00,
		},
		"P": {
			"Patricia Viana": 4000.00,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {

		fmt.Println(letra, funcs)
	}

}
