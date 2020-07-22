package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/many", many)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "Cookie1",
		Value: "123912",
	})
	fmt.Fprintln(w, "First cookie is written")
}

func read(w http.ResponseWriter, req *http.Request) {

	c1, err := req.Cookie("Cookie1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}
	fmt.Fprintln(w, "Your cookie no 1 is ", c1)

	c2, err := req.Cookie("Cookie2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}
	fmt.Fprintln(w, "Your cookie no 2 is ", c2)

	c3, err := req.Cookie("Cookie3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}
	fmt.Fprintln(w, "Your cookie no 3 is ", c3)

}

func many(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "Cookie2",
		Value: "91278324",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "Cookie3",
		Value: "1234982",
	})
	fmt.Fprintln(w, "All your cookies are written")
}
