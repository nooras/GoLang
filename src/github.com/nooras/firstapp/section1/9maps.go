package main

import "fmt"

func main() {
	//ss := map[string]int{"A": 1, "B": 2, "C": 3} //Direct declaration
	ss := make(map[string]int) //ss := make(map[string]int,3)
	ss = map[string]int{"A": 1, "B": 2, "C": 3}
	fmt.Println(ss)
	fmt.Println(ss["A"])
	ss["D"] = 4
	fmt.Println(ss)
	delete(ss, "D")
	fmt.Println(ss)
	//pop, ok := ss["D"] //value,true(present) or false(absent)
	_, ok := ss["D"]
	fmt.Println(ok)
	fmt.Println(len(ss))
	s := ss
	delete(s, "A")
	fmt.Println(s)
	fmt.Println(ss) // Delete will be from parent as well
}
