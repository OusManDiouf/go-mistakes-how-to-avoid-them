package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {

	rootHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		err := foo(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return // always return after an error in a http handler
		}
		_, _ = fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	s := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 500 * time.Millisecond,
		ReadTimeout:       500 * time.Millisecond,
		Handler:           http.TimeoutHandler(rootHandler, time.Second, "Server Timeout"),
		/**
		http.TimeoutHandler wrap the rootHandler and if the rootHandler fail to respond
		within 1s, the server return a 503 status code with "Server Timeout" as http response
		*/
	}

	log.Fatal(s.ListenAndServe())
}

func foo(r *http.Request) error {
	return nil
}
