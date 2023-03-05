package main

import "fmt"

func main() {

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

	//_ = memoryHogChanMerge(ch1, ch2)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- fmt.Sprintf("ch1=%d", i)
		}
		close(ch1)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- fmt.Sprintf("ch2=%d", i*10)
		}
		close(ch2)
	}()

	ch3 := make(chan string, 1)
	go func() {
		for i := 0; i < 10; i++ {
			ch3 <- fmt.Sprintf("ch3=%d", i*1000)
		}
		close(ch3)
	}()

	ch4 := make(chan string, 1)
	go func() {
		for i := 0; i < 10; i++ {
			ch4 <- fmt.Sprintf("ch4=%d", i*10000)
		}
		close(ch4)
	}()

	ch := ChanMerge(ch1, ch2, ch3, ch4)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		close(ch)
	}()

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
			for checkChans(chs...) {
				select {
				case v, open := <-chi: // this case is removed from the select statement once chi in nil
					if !open {
						chi = nil
						fmt.Println(fmt.Sprintf("ch%d closed", i))
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

func memoryHogChanMerge(ch1, ch2 chan int) chan int {
	/*
		There is one major issue: when one of the two channels is closed,
		the for loop will act as a busy-waiting loop,
		meaning it will keep looping even though no new message
		is received in the other channel.
	*/
	ch := make(chan int, 1)
	ch1Closed := false
	ch2Closed := false

	go func() {
		for {
			select {
			case v, open := <-ch1:
				if !open {
					ch1Closed = true
				}
				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2Closed = true
				}
				ch <- v
			}
			if ch1Closed && ch2Closed {
				close(ch)
				return
			}
		}
	}()
	return ch
}
