package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", serve)

	http.ListenAndServe(":80", nil)
}

func serve(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my Go webserver"))
}
