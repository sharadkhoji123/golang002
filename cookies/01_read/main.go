package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w,
		&http.Cookie{
			Name:  "MyFirst-Cookie",
			Value: "1234832-A",
		})

	fmt.Fprintln(w, "Cookie Written- Check your browser")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("MyFirst-Cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}
	fmt.Fprintln(w, "Your cookie ", c)
}
