package main

import (
	"fmt"
)

func main() {
	var x bool = true
	fmt.Printf("%v,%T\n", x, x)

	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	var y bool
	fmt.Printf("%v, %T\n", y, y) //Bydefault false value
}
