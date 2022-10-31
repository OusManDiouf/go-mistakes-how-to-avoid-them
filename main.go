package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	ch := make(chan Event)
	go consumer(ch)

	// initialise un chrono
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	// notif channel
	done := make(chan bool)
	// notifie qu'on à terminé
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	// listerners
	for {
		select {
		case <-done: // quitte la boucle au bout de 10s
			fmt.Println("Done!")
			return
		case t := <-ticker.C: // processing
			ch <- Event{
				message:   fmt.Sprintf("message%v", t.Second()),
				createdAt: t,
			}
		}
	}
}

type Event struct {
	message   string
	createdAt time.Time
}

func consumer(ch <-chan Event) {

	timerDuration := 5 * time.Second
	// NewTimer renvoie le currentTime sur son channel aprés timeDuration
	// ce qui permet au listerner (notre select) d'effectuer l'action necessaire
	// en cas de reception via le channel du timer
	// ici, il log un warning comme quoi aucune action voulue n'a été effectuer
	// dans la periodes de timeDuration(5s)
	// le timer est remis à zero avec une durée à chaque début de loop
	// le for loop traite les cases du select: pour chaque case une iteration...
	//
	timer := time.NewTimer(timerDuration)
	for {
		timer.Reset(timerDuration)
		select {
		case e := <-ch:
			fmt.Println(e)
		case <-timer.C:
			log.Printf("warning: no message received after %v seconds\n", timerDuration)
		}
	}
}
