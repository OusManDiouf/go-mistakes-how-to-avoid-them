package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := sync.WaitGroup{}
	var v uint64
	for i := 0; i < 30000; i++ {
		/**
		Le data race ici ce justifie par le fait qu'on ai invoqué wg.Add(1)
		dans la goRoutine, du coup, il n'y a aucune garantie qu'on ait indiqué
		au Wait qu'on veut bel et bien wait 3 goRoutines avant d'invoquer wg.Wait()
		SOLUCE: Add doit être invoqué avant d'éxecuter la goRoutine dans la goROutine parent
		et si le nbre de goRoutine est connu d'avance , on peut extraire Add completement de la boucle.
		Et Done doit être invoqué dans la goRoutine.
		*/
		wg.Add(1)
		go func() {
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v)
}
