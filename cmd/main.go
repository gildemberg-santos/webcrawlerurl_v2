package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
