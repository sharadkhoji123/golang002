package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	//	http.HandleFunc("/dog.jpg", dogP)

	http.Handle("/public/pics/", http.StripPrefix("/public/pics/", http.FileServer(http.Dir("./public/pics"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="public/pics/dog.jpeg">`)
}

// func dogP(w http.ResponseWriter, req *http.Request) {
// 	http.ServeFile(w, req, "public/pics/dog.jpg")
// }
