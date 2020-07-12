package goarea

import "math"

//Pi é o valor de PI
const Pi = 3.1416

//Circ calcula a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect calcula area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é uma função pública
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
