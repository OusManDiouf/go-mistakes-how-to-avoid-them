package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	var client *http.Client

	if getTracing() {
		client, err := createClientWithTracing()
		if err != nil {
			return
		}
		log.Println("withTracing: ", client)
	} else {
		client, err := createDefaultClient()
		if err != nil {
			return
		}
		log.Println("Default: ", client)
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
