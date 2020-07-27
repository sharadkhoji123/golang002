package main

import (
	"fmt"
	"net/http"
)

func check(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Fprintln(w, `<h1>Internal Server Error</h1>`)
	}
}
