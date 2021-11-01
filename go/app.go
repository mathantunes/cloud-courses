package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello")

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "Hi there")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}
