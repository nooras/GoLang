//Pointers is just a variable it stores the address of another variable
package main

import "fmt"

//Swaping address of 2 variable
func swap(m1,m2 *int){ //Both are type of int
	var temp int
	temp = *m2 //Storing address of temp
	*m2 = *m1  
	*m1 = temp
	
}

func main(){
	m1,m2 := 2,3
	fmt.Println(m1,m2)
	swap(&m1,&m2)
	fmt.Println(m1,m2)
	//ptr := &m1  //Stores the address of m1. & is refencing symbol also called refernece operators 
	//fmt.Println(ptr,*ptr) //For derefencing the address * is used 
}