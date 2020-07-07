package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	//fmt.Println("O valor de x é" + x) <---- nao permite concatenação direta entre String e int

	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs)

	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f", x) //printa formatado igual ao C (%.2f = float com duas casas decimais)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d) //%v para tipos genéricos, porém recomenda-se utilizar o tipo correto
}
