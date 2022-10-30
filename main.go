package main

func main() {

	// this channel is nil. the main goRoutine won't panic. It will just block FaEver !
	//var ch chan int
	//<-ch

	// the same apply if we send through a nil chan. the main goRoutine won't panic. It will just block FaEver !
	//var ch chan int
	//ch <- 0

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	_ = memoryHogChanMerge(ch1, ch2)

}

func memoryHogChanMerge(ch1, ch2 chan int) <-chan int {
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
