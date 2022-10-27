package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	for i := 0; i < 10; i++ {
		nonDeterministicExampleOfAntiDataRace()
	}
	//antiDataRaceV3()
	//antiDataRaceV2()
	///antiDataRaceV1()
}

func nonDeterministicExampleOfAntiDataRace() {
	i := 0
	mu := sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(2)

	func() {
		mu.Lock()
		defer wg.Done()
		defer mu.Unlock()
		i = 111
	}()

	go func() {
		mu.Lock()
		defer wg.Done()
		defer mu.Unlock()
		i = 222
	}()
	wg.Wait()
	fmt.Println("i: ", i)

}

func antiDataRaceV3() {
	i := 0
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	go func() {
		ch <- 1
	}()
	i += <-ch
	i += <-ch

	fmt.Println("i:", i)
}

func antiDataRaceV2() {
	i := 0
	mu := sync.Mutex{}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		mu.Lock()
		i++
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		mu.Lock()
		i++
		mu.Unlock()
	}()

	wg.Wait()
	fmt.Println(i)
}
func antiDataRaceV1() {
	var i int64

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		wg.Done()
		atomic.AddInt64(&i, 1)
	}()
	go func() {
		wg.Done()
		atomic.AddInt64(&i, 1)
	}()

	wg.Wait()
	fmt.Println(i)
}
