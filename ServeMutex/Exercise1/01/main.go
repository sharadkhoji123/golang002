package main

import (
	"io"
	"net/http"
)

type hotdog int

func d(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "Wuff Wuff")
}

func m(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "My name is John Ciena")
}

func b(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is an Index Page")
}
func main() {

	http.HandleFunc("/", b)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", m)
	http.ListenAndServe(":8080", nil)
}
