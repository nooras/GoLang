package main

import (
	"fmt"
)

//Unfuffered channel
func main(){
	//var c chan int 
	//fmt.Println(c) //Prints <nil>

	c := make(chan int) //Allocating memory

	//Send in a goroutine 
	go func(){
		c <- 1  //arrow means putting value into the channel and value sholud be same datatype
	}()

	//sniff
	val := <-c  //That means retriving data from channel and put it in val
	fmt.Println(val)
	fmt.Println(c)  //0xc000094060
}