package main

import "fmt"

func main() {

	a1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println("a1: ", a1)

	s1 := a1[0:5]

	fmt.Println("s1: ", s1)

	s1[3] = 99

	fmt.Printf("a1: %d\ns1: %d", a1, s1)

	s2 := s1[:4]
	fmt.Println("\ns2: ", s2)

	s2[3] = 50

	fmt.Println("s2: ", s2)
	fmt.Println("s1: ", s1)
}
