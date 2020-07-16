package main

import (
	"fmt"
)

type Car struct{
	Model string
}

//Buffer Channel
func main(){
	c := make(chan int, 3)
	d := make(chan *Car, 3)
	go func(){
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		close(c)  //It indicates there is no channel
		d <- &Car{"1"}
		d <- &Car{"2"}
		d <- &Car{"3"}
		d <- &Car{"4"}
		close(d)
	}()

	for i := range c {
		fmt.Println(i)
	}

	for i := range d {
		fmt.Println(i.Model)
	}
}