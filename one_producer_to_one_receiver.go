package main

import "fmt"

type handle struct {
	value          int
	isDisconnected bool
}

func main() {
	/*
		--------------------
		ONE(producer) TO ONE(receiver) RELATION
		--------------------
		// En utilisant un "unbuffered Channel" ici, on est garantie que le producer (main)
		// va bloquer jusqu'à ce que le receiver (select first case) soit prêt,
		// cette approche garantie que tous les messages vers messageCh (unbuffchan)
		// seront bien receptionné dans le bon ordre !!
	*/
	// ON CHANNEL HANDLER (use unbeffered channel)
	messageCh := make(chan handle)

	go func() {
		for {
			v := <-messageCh
			if v.isDisconnected == false {
				fmt.Println(v.value)
			}
			if v.isDisconnected == true {
				fmt.Println("disconnection, return")
				return
			}
		}
	}()

	// SINGLE PRODUCER
	for i := 0; i < 2_000_000; i++ {
		message := handle{
			value:          i,
			isDisconnected: false,
		}
		messageCh <- message
	}
	// send a disconnect message
	messageCh <- handle{
		value:          0,
		isDisconnected: true,
	}
}
