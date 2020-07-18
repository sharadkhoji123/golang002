package main

import (
	"fmt"
	"net/http"
)

type hotdog struct {
	val int
}

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "And the value is := ", m.val)
}

func main() {

	t := hotdog{
		val: 5,
	}

	http.ListenAndServe(":8080", t)

}
