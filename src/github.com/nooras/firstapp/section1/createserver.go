package main

import(
	"net/http"
	"fmt"
)

func main(){
	mux := http.NewServeMux()
	//Router defines an end point /  if / will hit then run this function will executed
	mux.HandleFunc("/",func(w http.ResponseWriter  , r *http.Request){
		fmt.Println("Request Received")
		fmt.Println(r.Method)
		w.Write([]byte("Hello World"))
	})
	mux.HandleFunc("/go",func(w http.ResponseWriter  , r *http.Request){
		fmt.Println("Request Received")
		fmt.Println(r.Method)
		w.Write([]byte("Hello Nooras Fatima"))
	})
	http.ListenAndServe("localhost:1234",mux) //Creating server which is listening at this port and passed multiplexer and router 
}