//Goroutines : It allows methods to run in background without waiting for new funtion.
//Allows executing the next statement without waiting previous statement 

package main

import (
	"fmt"
	"time"
)

func heavy(){
	for{
		time.Sleep(time.Second * 1)
		fmt.Println("Hello in heavy")
	}
}

func superHeavy(){
	for{
		time.Sleep(time.Second * 2)
		fmt.Println("Hello in super heavy")
	}	
}

func main(){
	go heavy()  //Asynchrounous function
	go superHeavy()  //Both function will executeing concurrently
	fmt.Println("Executed") 
	select { } //Indefinately run
}