package main

import "fmt"

func notaParaConceito(nota float64) string {
	n := int(nota)
	switch {

	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	case n >= 0 && n < 3:
		return "E"
	default:
		return "Nota inválida"

	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(8.1))
	fmt.Println(notaParaConceito(5.3))
	fmt.Println(notaParaConceito(3.3))
	fmt.Println(notaParaConceito(1.2))
	fmt.Println(notaParaConceito(-2))
}
