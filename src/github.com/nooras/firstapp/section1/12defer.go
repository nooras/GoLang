package main

import "fmt"

//Defer
//Run at the end of main function
//Last in first out runs
func main() {
	fmt.Println("hii")
	defer fmt.Println("Good")
	panic("Something bad")
	defer fmt.Println("Morn!")
}
