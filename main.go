package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	address := "localhost:9000"
	fmt.Println("Address start at:", address)
	start := http.ListenAndServe(address, nil)
	if start != nil {
		fmt.Println(start.Error())
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message string = "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message string = "Hello world!"
	w.Write([]byte(message))
}
