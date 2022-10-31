package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu       sync.Mutex
	counters map[string]int
}

//NewCounter is a factory function
func NewCounter() Counter {
	return Counter{
		counters: map[string]int{},
	}
}

func (c *Counter) Increment(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	counter := NewCounter()
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
