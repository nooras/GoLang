package main

import (
	"fmt"
)

func main() {
	a := 8              //2^3
	fmt.Println(a >> 3) //1   bit shift right - 2^3 * 2^3 = 2^6
	fmt.Println(a << 3) //64  bit shift left - 2^3 / 2^3 = 2^0
}
