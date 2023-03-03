package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Taken from kennedy
//func init() {
//	// here we are running against 2P
//	runtime.GOMAXPROCS(2)
//	//rand.Seed(time.Now().UnixNano())
//}

func main() {

	/**
	This 3 operations will eventually do they job
	regardless of the execution order of the there goroutines
	There is no data-race
	But we still have a race condition issue.
	*/
	//DataRaceSolutionAtomic()
	//DataRaceSolutionMutex()
	//DataRaceSolutionChannel()

	//for i := 0; i < 100_000; i++ {
	//	NonDeterministicCode()
	//}

	//Solution : S'assurer que les deux goroutines se lance dans le bon ordre.
	for i := 0; i < 100_000; i++ {
		DeterministicCodeUsingChannel()
	}

	//fmt.Println(runtime.GOOS)
	//fmt.Println(runtime.Compiler)
	//fmt.Println(runtime.GOARCH)
	//fmt.Println(runtime.Version())
}

// NonDeterministicCode
//	but data race free !
func NonDeterministicCode() {
	// Il n'y a pas de data-race ici,
	// mais on est bien en presence d'une race condition,
	// du fait de l'aspect non-deterministic :
	// le résultat depend de la sequence ou le timing des events
	// qui ne peuvent pas être contrôlé.
	// Here, the timing of events is the goroutines’ execution order.

	i := 0
	mu := sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		mu.Lock()
		defer mu.Unlock()
		i += 2
	}()

	func() {
		defer wg.Done()

		mu.Lock()
		defer mu.Unlock()
		i *= 2_000_000
	}()

	wg.Wait()
	fmt.Println("I = ", i)

}

func DeterministicCodeUsingChannel() {
	i := 0
	ch := make(chan struct{})

	go func() {
		i += 2
		<-ch
	}()

	ch <- struct{}{}

	go func() {
		i *= 2_000_000
		<-ch
	}()

	ch <- struct{}{}
	fmt.Println("I = ", i)
}

// DataRaceSolutionChannel (RECOMENDED)
//		Share memory by communicating; don't communicate by sharing memory.
//      -> Use communication and channels to ensure that
//         a variable is updated by only one goroutine
func DataRaceSolutionChannel() {
	i := 0

	// Unbuffered channel: G sender wait until the G receiver is ready
	ch := make(chan int)
	go func() {
		ch <- 1 // "send" happen here
	}()
	go func() {
		ch <- 1 // "send" happen here
	}()

	// No need for a wg here, No data-race because,
	// the above senders will block until their respective receiver bellow are ready
	fmt.Println("I = ", i) // read happen

	// La main goroutine est la seule à accéder en ecriture à cette variable
	// cette fonction est donc free of data-race data-race !
	// À noter aussi que les deux ecritures se deroule de maniere sequentiel
	i += <-ch // G receiver (main routine) is ready here
	i += <-ch // G receiver (main routine) is ready here

	// No need for a wg here,No data-race,
	// because the "write" already happen above
	fmt.Println("I = ", i) // read happen

}

func DataRaceSolutionMutex() {
	i := 0
	mu := sync.Mutex{}

	var wg sync.WaitGroup
	wg.Add(2)

	// Parallel exec des deux goroutines
	// donc on a besoin d'un mécanisme
	// pour syncroniser l'accès à la shared rss i.
	go func() {
		defer wg.Done()
		mu.Lock()
		i++ // here is the critical section
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		mu.Lock()
		i++ // here is the critical section
		mu.Unlock()
	}()

	wg.Wait()

	// Cette instruction peut induire un data-race
	// à cause du syscall qu'entraine Println
	// c'est ce qui justifie l'usage du wg
	fmt.Println("I = ", i)
}
func DataRaceSolutionAtomic() {
	var i int64

	var wg sync.WaitGroup
	wg.Add(2)

	// Parallel exec des deux goroutines
	// donc on a besoin d'un mécanisme
	// pour syncroniser l'accès à la shared rss i.
	go func() {
		defer wg.Done()
		atomic.AddInt64(&i, 1)
	}()
	go func() {
		defer wg.Done()
		atomic.AddInt64(&i, 1)
	}()

	wg.Wait()
	// Cette instruction peut induire un data-race
	// à cause du syscall qu'entraine Println
	// c'est ce qui justifie l'usage du wg
	fmt.Println("I = ", i)
}
