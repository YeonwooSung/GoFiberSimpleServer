package main

import (
	"fmt"
	"net/http"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Bar!")
}

func main() {
	fmt.Println("Init LinkPad server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
		//TODO
	})

	http.HandleFunc("/bar", barHandler)

	http.ListenAndServe(":3000", nil)
}
