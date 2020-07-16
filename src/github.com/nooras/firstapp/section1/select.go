 package main

 import (
	 "fmt"
	"os"
	"time"
)

func Select(c chan int, quitz chan struct{}){
	for{
		time.Sleep(time.Second)
		select{
		case <- c:
			fmt.Println("Recieved")
		case <- quitz:
			fmt.Println("Quit")
			os.Exit(0)
		}
	}
}

func main(){
	c := make(chan int, 1)
	quitz := make(chan struct{})
	go func(){
		c <- 1
		quitz <- struct{}{}
	}()

	go Select(c,quitz)
	select {}
}