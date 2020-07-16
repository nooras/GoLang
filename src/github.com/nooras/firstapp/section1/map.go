package main

import "fmt"

func main(){
	mymap := make(map[string]interface{})

	mymap["name"] = "Nooras"
	mymap["age"] = 12
	fmt.Println(mymap,mymap["name"])
}