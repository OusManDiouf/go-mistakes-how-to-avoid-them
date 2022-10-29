package main

import (
	"context"
	"fmt"
	"time"
)

type Position struct {
	lat  int
	long int
}

type publisher interface {
	Publish(ctx context.Context, ch chan Position)
}

type radarPublisher struct{}

func (rp radarPublisher) Publish(ctx context.Context, ch chan Position) {
	for {
		select {
		case v := <-ch:
			fmt.Printf("%v sent to kafka\n", v)
			break
		case <-ctx.Done():
			fmt.Println("Timeout.")
			return
		}
	}
	//time.Sleep(10 * time.Second)
	//fmt.Println("Position published to kafka...")
}

type publishHandler struct {
	pub publisher
}

func (h publishHandler) publishPosition(p Position) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	ch := make(chan Position)
	go func() {
		time.Sleep(5 * time.Second)
		ch <- p
	}()
	h.pub.Publish(ctx, ch)
	return nil
}

func main() {
	position := Position{
		lat:  10,
		long: 20,
	}
	radarPub := radarPublisher{}

	h := publishHandler{pub: radarPub}
	err := h.publishPosition(position)
	if err != nil {
		return
	}
}
