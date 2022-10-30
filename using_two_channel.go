package main

import "fmt"

func main() {

	// En utilisant un "unbuffered Channel" ici, on est garantie que le producer (main)
	// va bloqué jusqu'a ce que le receiver (processMessage) soit prêt,
	// cette approche garantie que tout les messages vers messageCh (unbuffchan)
	// seront bien receptionné avec que celle du disconnectCh !!
	messageCh := make(chan int)
	disconnectCh := make(chan struct{})

	go processMessagesTwoCh(messageCh, disconnectCh)

	for i := 0; i < 5; i++ {
		messageCh <- i
	}
	disconnectCh <- struct{}{}
}
func processMessagesTwoCh(messageCh <-chan int, disconnectCh <-chan struct{}) {
	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			fmt.Println("disconnection, return")
			return
		}
	}
}
