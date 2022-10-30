package main

import (
	"fmt"
)

/**
	Dans le cas de figure où l'on a plusieurs producers, il peut être pratiquement impossible
    de garantir l'ordre d'écriture.
    Qu'on utilise un UNBUFFERED_CHANNEL ou un SINGLE_CHANNEL, cela nous menera inevitablement
	à un race condition !!
	SOLUTION:
		-> Dans le cas où l'on a plusieurs receiver
         	-> Recevoir à partir de messageCh ou disconnectCh
			-> Si disconnectCh reçoit
				-> Lire tout les messages de messageCh, s'ils existent
				-> Puis return
*/
func main() {
	// En utilisant un "unbuffered Channel" ici, on est garantie que le producer (main)
	// va bloqué jusqu'a ce que le receiver (processMessage) soit prêt,
	// cette approche garantie que tout les messages vers messageCh (unbuffchan)
	// seront bien receptionné avec que celle du disconnectCh !!
	messageCh := make(chan int)
	disconnectCh := make(chan struct{})

	go processMessage(messageCh, disconnectCh)

	for i := 0; i < 5; i++ {
		go func(message int) {
			messageCh <- message
		}(i)
	}
	disconnectCh <- struct{}{}

}

func processMessage(messageCh <-chan int, disconnectCh <-chan struct{}) {

	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			fmt.Println("disconnect reached before endind messageCh procession...continuing")
			for {
				select {
				case v := <-messageCh:
					fmt.Println(v)
				default:
					fmt.Println("disconnection, return")
					return
				}

			}
		}
	}
}
