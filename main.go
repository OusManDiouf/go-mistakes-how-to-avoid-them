package main

import (
	"math/rand"
	"net/http"
	"time"
)

func main() {

	var client *http.Client
	var err error

	if getTracing() {
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}
	// mutualize error handling
	if err != nil {
		return
	}
	println(client)
}

func createClientWithTracing() (*http.Client, error) {
	return &http.Client{}, nil
}
func createDefaultClient() (*http.Client, error) {
	return &http.Client{
		Timeout: time.Second,
	}, nil
}

func getTracing() bool {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(2)
	switch r {
	case 1, 3:
		return true
	default:
		return false
	}
}
