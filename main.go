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

	http.HandleFunc("/users", func(wr http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// HTTP GET
			fmt.Println("GET")
		case http.MethodPost:
			// HTTP post
			fmt.Println("POST")
		}
	})

	http.HandleFunc("/bar", barHandler)

	http.ListenAndServe(":8080", nil)
}
