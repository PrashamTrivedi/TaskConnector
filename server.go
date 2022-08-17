package main

import (
	"fmt"
	"net/http"
	"taskConnector/cmd"
)

func StartServer() {
	mux := defaultMux()
	http.ListenAndServe(":5000", mux)
	fmt.Println("Starting the server: http://localhost:5000")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}
func hello(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("key")
	if path != "" {
		output := cmd.RunCommand(path)
		fmt.Fprintln(w, output)
	} else {
		fmt.Fprint(w, "Hello there")
	}
}
