package main

import (
	"fmt"
)

/**

	--------------------
	MANY(producers) TO MANY(receivers) RELATION
	--------------------
	Dans le cas de figure où l'on a plusieurs producers,
    il peut être pratiquement impossible de garantir l'ordre d'écriture.
    Qu'on utilise un UNBUFFERED_CHANNEL ou un SINGLE_CHANNEL,
    cela nous menera inevitablement à un "race" condition !!
	SOLUTION:
		-> Dans le cas où l'on a plusieurs receiver
         	→ Recevoir à partir de messageCh ou disconnectCh
			→ Si disconnectCh reçoit
				→ Lire tous les messages de messageCh, s'ils existent
				→ Puis return
*/
func main() {
	// En utilisant un "unbuffered Channel" ici, on est garantie que le producer (main)
	// va bloquer jusqu'à ce que le receiver (processMessage) soit prêt,
	// cette approche garantie que tous les messages vers messageCh (unbuffchan)
	// seront bien receptionné avec que celle du disconnectCh !!

	// MANY CHANNELS RECEIVER
	messageCh := make(chan int)
	disconnectCh := make(chan struct{})

	go processMessage(messageCh, disconnectCh)

	// MANY PRODUCERS
	for i := 0; i < 10; i++ {
		go func(i int) {

			// we are missing some message when we use an async sender
			// It's because some messageCh are sent after the goroutine has returned.
			// messageCh <- sender.GetMessage(i)
			messageCh <- i
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
			fmt.Println("disconnect reached...processing remaining message before disconnecting")
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
