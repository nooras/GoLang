package main

import (
	"fmt"
	"sync"
)

func main(){
	// func(){
	// 	fmt.Println("Hello")
	// }()
	// func(){
	// 	fmt.Println("World")
	// }()
	// fmt.Println("!!")

	//wait groups
	wg := &sync.WaitGroup{}
	//add , done and wait

	//mutexes
	wg.Add(2)
	go func(){
		fmt.Println("Hello")
		wg.Done()
	}()
	go func(){
		fmt.Println("World")
		wg.Done()
	}()
	fmt.Println("!!")
	wg.Wait()
	fmt.Println("$$")
}