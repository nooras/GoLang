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

func main(){
	go heavy()  //Asynchrounous function
	fmt.Println("Executed") 
	time.Sleep(time.Second * 5)  //BAckground function will run in background and it will show it on screen
}