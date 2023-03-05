package main

import (
	"fmt"
	"github.com/OusManDiouf/go-mistakes-how-avoid-them/sender"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	/*
		--------------------
		ONE(producer) TO MANY(receivers) RELATION
		--------------------
		En utilisant un "unbuffered Channel" ici, on est garantie que le producer (main)
		va bloquer jusqu'à ce que le receiver (the first case) soit prêt,
		cette approche garantie que tous les messages vers messageCh (unbuffchan)
		seront bien receptionné avec que celle du disconnectCh !!
	*/
	// MANY CHANNELS RECEIVERS (use unbuferred channel)
	messageCh := make(chan int)
	disconnectCh := make(chan struct{})

	go func() {
		for {
			select {
			case v := <-messageCh:
				fmt.Println(v)
			case <-disconnectCh:
				fmt.Println("disconnection, return")
				return
			}
		}
	}()

	// SINGLE PRODUCER
	for i := 0; i < 100; i++ {
		messageCh <- sender.GetMessage(i) // The G sender block until the G receiver is ready
	}
	disconnectCh <- struct{}{}
}
