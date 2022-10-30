package main

import "fmt"

type handle struct {
	value          int
	isDisconnected bool
}

func main() {

	// En utilisant un "unbuffered Channel" ici, on est garantie que le producer (main)
	// va bloqué jusqu'a ce que le receiver (processMessage) soit prêt,
	// cette approche garantie que tout les messages vers messageCh (unbuffchan)
	// seront bien receptionné avec que celle du disconnectCh !!
	messageCh := make(chan handle)
	//disconnectCh := make(chan struct{})

	go processMessagesOneCh(messageCh)

	for i := 0; i < 2_000_000; i++ {
		message := handle{
			value:          i,
			isDisconnected: false,
		}
		messageCh <- message
	}
	messageCh <- handle{
		value:          0,
		isDisconnected: true,
	}
}
func processMessagesOneCh(messageCh <-chan handle) {
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
}
