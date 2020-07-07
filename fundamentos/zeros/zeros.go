package main

import "fmt"

func main() {
	var (
		a int
		b float64
		c bool
		d string
		e *int
	)

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)

}
