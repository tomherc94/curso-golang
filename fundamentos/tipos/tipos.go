package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numeors inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	/*
		sem sinal (só positivos)... uint8 uint16 uint32 uint64 (equivalente a byte, short, int e long)
		u = unsigned
	*/

	var b byte = 255 //byte é um alias para uint8
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64 (equivalente a byte, short, int e long)

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	//numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s1 := "Olá meu nome é Tomás"
	fmt.Println(s1 + "!")
	fmt.Println("O tipo de s1 é", reflect.TypeOf(s1))
	fmt.Println("O tamanho de s1 é", len(s1))

	//string com multiplas linhas (usa-se crase)
	s2 := `Olá
	meu
	nome
	é
	Tomás
	!`

	fmt.Println("O tamanho de s2 é", len(s2))

	//char == int32
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))

}
