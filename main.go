package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type handler struct {
	client http.Client
	url    string
}

func (h handler) GetBody() (string, error) {

	res, err := h.client.Get(h.url)
	if err != nil {
		log.Fatalf("failed to request the url %v: %v\n", h.url, err)
		return "", err
	}
	// Do not forget to close the body rss: res return a pointer on http.Response
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Fatalf("failed to close response: %v\n", err)
			return
		}
	}()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("failed to get body from response: %v\n", err)
		return "", err
	}

	return string(body), nil
}
func main() {

	h := handler{
		client: http.Client{},
		url:    "https://api.chucknorris.io/jokes/random",
	}

	body, err := h.GetBody()
	if err != nil {
		return
	}

	fmt.Println(body)

}
