package main

import (
	"fmt"
	"net/http"
)

func StartServer() {
	mux := defaultMux()
	http.ListenAndServe("5000", mux)
	fmt.Println("Starting the server: http://localhost:5000")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
