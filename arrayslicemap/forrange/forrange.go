package main

import "fmt"

func main() {

	//semelhante ao for each!!

	numeros := [...]int{1, 2, 3, 4, 5} //compilador conta o tamanho do array

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}

}
