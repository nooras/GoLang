package main

import "fmt"

//int,int8, int16, int32,int64

func main() {

	var m1 int //single declaration
	m1 = 2
	fmt.Println(m1)

	var ( //Multiple declaration
		v1 = 1 //automatically implicit int
		v2 = 2
	)
	fmt.Println(v1 + v2)

	var c1 int
	var c2 int32                //cant add 2 diffent datatypes
	fmt.Println(int32(c1) + c2) //Bydefault it takes 0 values

	a1 := 5 //Direct Inititalization andsubstitution
	a2 := 5
	fmt.Println(a1 + a2)

	var x int = 10
	fmt.Println(x)
	fmt.Printf("%v,%T\n", x, x) //Value, Type

	var y float32
	y = float32(x) //Conversion
	fmt.Printf("%v,%T\n", y, y)
}
