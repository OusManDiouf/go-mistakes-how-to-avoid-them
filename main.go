package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	// Here s is evaluated only once!
	// s is copied in mem and the looping occur on this copy
	// hence, even that we keep appending to the original slice,
	// the range operate on the copy, which keep the same length
	// that's why the loop stop after 3 iterations.
	for range s {
		s = append(s, 10)
	}
	fmt.Println(s)

	ch1 := make(chan int, 3)
	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		close(ch1)
	}()

	ch2 := make(chan int, 3)
	go func() {
		ch2 <- 10
		ch2 <- 20
		ch2 <- 30
		close(ch2)
	}()

	ch := ch1
	// remember, here, ch(pointing to ch1) is only evaluated once !
	// and is copied in mem for the looping
	for v := range ch { // the range create a channel consumer
		// Here we consume 3 values from ch1
		fmt.Println(v)
		// no effect on the current loop, the copy of ch is already used for the looping
		// and here we already jump in the loop.
		// but the assignement take effect, if the channel is closed from the sender side
		// it will close ch2!
		ch = ch2
	}
	// here we can proove that ch never get closed due to the above reassignment
	_, okCh := <-ch
	if !okCh {
		fmt.Println("ch closed")
	} else {
		fmt.Println("ch not closed yet")
	}
	// here ch1 is closed
	_, okCh1 := <-ch1
	if !okCh1 {
		fmt.Println("ch1 closed")
	} else {
		fmt.Println("ch1 not closed yet")
	}
	// here ch2 is not closed
	_, okCh2 := <-ch2
	if !okCh2 {
		fmt.Println("ch2 closed")
	} else {
		fmt.Println("ch2 not closed yet")
	}

	// and outside the loop ch is now pointing to ch2
	// here we are iterating over the element of ch2 !
	//for v := range ch {
	//	fmt.Println(v)
	//}

}
