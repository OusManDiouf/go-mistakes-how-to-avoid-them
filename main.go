package main

import (
	"fmt"
	"sync"
	"time"
)

type Donation struct {
	cond    *sync.Cond
	balance int
}

func main() {

	donation := Donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// Listene goRoutines
	listenerFunc := func(goal int) {
		donation.cond.L.Lock()
		// if the goal is not met yet
		// First Critical section
		for donation.balance < goal {
			// wait for a condition (balance update) within lock/unlock
			// wait marche comme suit:
			//	-> Unlock le mutex
			//	-> Suspend la goRoutine courrante, et atteint une notification
			// -> Lock le mutex lorqu'une notification arrive
			donation.cond.Wait()
		}
		// Second Critical section
		fmt.Printf("%d goal reached\n", donation.balance)
		donation.cond.L.Unlock()
	}

	go listenerFunc(10)
	go listenerFunc(15)

	// Updater goRoutine
	for {
		time.Sleep(time.Second)
		fmt.Println("Current Balance: ", donation.balance)
		donation.cond.L.Lock()
		// Critical section: locked to prevent data races
		donation.balance++
		donation.cond.L.Unlock()
		// reveille tous les goRoutines qui attende sa condition (à savoir à chaque update)
		donation.cond.Broadcast()
	}
}
