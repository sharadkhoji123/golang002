package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/Cute_dog.jpg", cuteDog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, `
	<img src="Cute_dog.jpg">
	`)
}

func cuteDog(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("Cute_dog.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
	}
	defer f.Close()
	io.Copy(w, f)
}
