package main

import "fmt"

//Interfaces are name collection of methods signatures

//Interface
type Car interface{
	Drive()
	Stop()
}

//For confirmaing all interface is used or not
//Follow same instruction of interface or not
func NewModel(arg string) Car{
	return &Lambo{arg}
}

type Lambo struct{
	LamboModel string
}

type Chevy struct{
	ChevyModel string
}

func (l *Lambo) Drive(){
	fmt.Println("Lambo on the move")
	fmt.Println(l.LamboModel)
}

func (l *Lambo) Stop(){
	fmt.Println("Stopping lambo.")
	fmt.Println(l.LamboModel)
}

func (c *Chevy) Drive(){
	fmt.Println("Chevy on the move")
	fmt.Println(c.ChevyModel)
}

func main(){
	//l := Lambo{"Gallardo"}
	l := NewModel("Gallardo")
	c := Chevy{"C369"}
	l.Drive()
	c.Drive() 
}