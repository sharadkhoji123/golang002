package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "Wuff Wuff Wuff!!!")
	case "/cat":
		io.WriteString(w, "Miaooo Miaooo Miaooo!!!")
	default:
		io.WriteString(w, "<h1>Welcome to the home page of The Dog and The Cat</h1>")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
