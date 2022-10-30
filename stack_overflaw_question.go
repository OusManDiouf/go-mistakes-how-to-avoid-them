package main

import (
	"fmt"
	"strconv"
	"strings"
)

type message struct {
	body string
	code int
}

var markets []string = []string{"BTC", "ETH", "LTC"}

// produces messages into the chan
func produce(n int, market string, msg chan<- message) {
	var msgToSend = message{
		body: strings.Join([]string{"market: ", market, ", #", strconv.Itoa(1)}, ""),
		code: 1,
	}
	fmt.Println("Producing:", msgToSend)
	msg <- msgToSend
}

func receive(msg <-chan message) {
	for {
		m, ok := <-msg
		if ok {
			fmt.Println("Received:", m)
		}
	}
}

func main() {
	msgC := make(chan message)

	go receive(msgC)

	for ix, market := range markets {
		go produce(ix+1, market, msgC)
	}
}
