package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) (string, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// ch := make(chan int)
	ch := make(chan string)
	go func() {
		// ridiculous long running task
		// uid := ctx.Value("userID").(int)
		uname := ctx.Value("fname").(string)
		time.Sleep(6 * time.Second)

		// check to make sure we're not running in vain
		// if ctx.Done() has
		if ctx.Err() != nil {
			return
		}
		ch <- uname
		// ch <- uid
	}()

	select {
	case <-ctx.Done():
		return "finish", ctx.Err()
	case i := <-ch:
		return i, nil
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
