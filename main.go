package main

import (
	"errors"
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := foo(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return // always return after an error in an http handler
		}
		_, _ = fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}))

	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}

func foo(r *http.Request) error {
	return errors.New("an error occured")
}
