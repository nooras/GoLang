package main

import "fmt"

//Cannot change
func main() {
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	const a int = 42
	var b int = 27
	fmt.Printf("%v,%T\n", a+b, a+b)

	//Multiple declaration iota start with 0 and increases by 1
	const (
		x = iota
		y = iota
		z = iota
	)
	const (
		x2 = iota
		y2
		z2
	)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", x2)
	fmt.Printf("%v\n", y2)
	fmt.Printf("%v\n", z2)

	const (
		//binary shift by 1 digit oit append zero to the right hand side
		isAdmin       = 1 << iota //1
		isHeadQuart               //2
		canSeeFunance             //2*2 ->2*2*2->2*2*2*2 and so on
	)
	fmt.Printf("%v\n", isAdmin)
	fmt.Printf("%v\n", isHeadQuart)
	fmt.Printf("%v\n", canSeeFunance)
}
