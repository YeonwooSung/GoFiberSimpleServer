package main

import (
	"fmt"
	"net/http"
)

/**
 * Error handling function to handle the 404 not found
 *
 */
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}

	//TODO handle more errors
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hello Bar!")
}

func main() {
	fmt.Println("Init LinkPad server")

	http.HandleFunc("/", indexHandler)
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

	http.ListenAndServe(":8080", nil)
}
