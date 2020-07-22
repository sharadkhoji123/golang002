package main

import (
	"fmt"
	"net/http"
	"time"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Mcleod-Key", "this is from mcleod")
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintln(w, "<h1>Write your code here at this time: </h1>", time.Now())
}

func main() {
	var t hotdog
	http.ListenAndServe(":8080", t)
}
