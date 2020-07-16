package main

import (
	"fmt"
)

func main() {
	var n complex64 = 1 + 2i //var n complex128=complex(1,5)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", real(n), real(n))
	fmt.Printf("%v, %T\n", imag(n), imag(n))
}
