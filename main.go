package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	// this channel is nil. the main goRoutine won't panic.
	// It will just block FaEver !
	//var ch chan int
	//<-ch

	// The same apply if we send through a nil chan.
	// The main goRoutine won't panic.
	// It will just block FaEver !
	//var ch chan int
	//ch <- 0

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)
	ch4 := make(chan string, 1)

	var wg sync.WaitGroup
	wg.Add(4)

	//_ = memoryHogChanMerge(ch1, ch2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			// simulating a long running task
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			ch1 <- fmt.Sprintf("ch[1][%d]", i)
		}
		close(ch1)
	}()
	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i++ {
			// simulating a long running task
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			ch2 <- fmt.Sprintf("ch[2][%d]", i)
		}
		close(ch2)
	}()

	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i++ {
			// simulating a long running task
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			ch3 <- fmt.Sprintf("ch[3][%d]", i)
		}
		close(ch3)
	}()

	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i++ {
			// simulating a long running task
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			ch4 <- fmt.Sprintf("ch[4][%d]", i)
		}
		close(ch4)
	}()

	// DRIVER GO ROUTINE - Collect the result of the merge channel
	ch := ChanMerge(ch1, ch2, ch3, ch4)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		close(ch)
	}()

	wg.Wait()
}

func ChanMerge(chs ...chan string) chan string {
	ch := make(chan string, 1)

	checkChans := func(chs ...chan string) bool {
		chanCount := len(chs)
		flag := 0
		for _, c := range chs {
			if c == nil {
				flag++
			}
		}
		if flag == chanCount {
			return false
		}
		return true
	}

	for i, chi := range chs {
		go func(i int, chi chan string) {
			i++
			for checkChans(chs...) {
				select {
				case v, open := <-chi: // this case is removed from the select statement once chi in nil
					if !open {
						chi = nil
						fmt.Println(fmt.Sprintf("ch[%d][X]", i))
						break
					}
					ch <- v
				}
			}
			close(ch)
		}(i, chi)
	}

	return ch
}
