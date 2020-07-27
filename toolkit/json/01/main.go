package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mrshl", mrshl)
	http.HandleFunc("/encde", encde)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>FOO</title>
		</head>
		<body>
		You are at foo
		</body>
		</html>`
	w.Write([]byte(s))
}

func mrshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	p1 := person{
		Fname: "Sharad",
		Lname: "Nautiyal",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}

	json, err := json.Marshal(p1)
	if err != nil {
		fmt.Fprintln(w, "There is an error while Marshalling..")
		return
	}
	w.Write(json)
}

func encde(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Don't", "Panic", "When", "Scared"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		fmt.Fprintln(w, "There is an error while Encoding...")
		return
	}
}
