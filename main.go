package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type key string

const tracingKey key = "tracingKey"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req := context.WithValue(r.Context(), tracingKey, "407576-8221")

		fmt.Printf("ctx in handler closed: %v\n", r.Context().Err() != nil)
		fmt.Printf("tracingKey in handler: %v\n", req.Value(tracingKey))

		// perform some task to compute the response
		response, err := doSomeTask(req, r)
		if err != nil {
			return
		}
		// create a goRoutine to publish the response to Kafka
		go func() {
			time.Sleep(1 * time.Second)
			// calling publish will return an error because we returned the Http response to the client quickly
			// the context associated with the request is canceled once we write the response to the client
			///publish(reqCtx, response)

			// SOLUTION: do not associate publish with the parent context, instead, create a new one.
			// Et si le context contient des valeurs utiles ?
			// Bien que ces valeurs soient bel et bien toujours dispo au niveau du context parent,
			// l'ideal serait d'avoir un tout nouveau context qui est détaché du mecanisme
			// d'annulation du contexte parent tout en incluant les valeurs utiles de ce ctx parent.
			// SOLUTION: implementer un custom context semblable à celui fourni par go
			//			 mais sans le signal d'annulation (cancellation signal)

			// create a new empty detached context, that wil never be closed or canceled,
			// but it will carry the parent context's Values
			publish(detach{ctx: req}, response)
		}()

		// write the http response
		_, err = w.Write([]byte("ok"))
		if err != nil {
			log.Fatalf("Error server %v\n", err)
			return
		}
	})

	log.Println("Server listening on localhost:4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatalf("Error server %v\n", err)
	}
}

func publish(ctx context.Context, response struct{}) {
	fmt.Printf("ctx in publish closed: %v\n", ctx.Err() != nil)
	fmt.Printf("tracingKey in publish: %v\n", ctx.Value(tracingKey))
}
func doSomeTask(ctx context.Context, r *http.Request) (struct{}, error) {
	fmt.Printf("ctx in doSomeTask closed: %v\n", ctx.Err() != nil)
	fmt.Printf("tracingKey in doSomeTask: %v\n", ctx.Value(tracingKey))
	return struct{}{}, nil
}
