package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	io.WriteString(w, "Hello "+vars["name"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", Hello)
	// http.Handle("/", r)
	http.ListenAndServe(":8000", r)
}
