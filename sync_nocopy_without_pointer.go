package main

import (
	"fmt"
	"sync"
)

type CounterV2 struct {
	mu       *sync.Mutex
	counters map[string]int
}

//NewCounterV2 is a factory function
func NewCounterV2() CounterV2 {
	return CounterV2{
		mu:       &sync.Mutex{},
		counters: map[string]int{},
	}
}

func (c CounterV2) Increment(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	counter := NewCounterV2()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		counter.Increment("foo")
	}()
	go func() {
		defer wg.Done()
		counter.Increment("bar")
	}()
	wg.Wait()
	fmt.Println(counter.counters)
}
