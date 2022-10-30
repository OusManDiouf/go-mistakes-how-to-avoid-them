package main

import (
	"fmt"
	"strconv"
	"sync"
)

type CacheV2 struct {
	//TODO: try with mu sync.Map
	mu       sync.RWMutex
	balances map[string]float64
}

func (c *CacheV2) AddBalance(id string, balance float64) {
	c.mu.Lock()
	c.balances[id] = balance
	c.mu.Unlock()
}

func (c *CacheV2) AverageBalance() float64 {
	c.mu.Lock()
	// soluce: if the iteration operation isn’t lightweight,
	// is to work on an actual copy of the data and protect only the copy
	m := make(map[string]float64, len(c.balances))
	for k, v := range c.balances {
		m[k] = v
	}
	defer c.mu.Unlock()

	sum := 0.
	for _, balance := range m {
		sum += balance
	}
	return sum / float64(len(m))
}

func (c *CacheV2) Size() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.balances)
}

func main() {
	cache := CacheV2{
		mu:       sync.RWMutex{},
		balances: map[string]float64{},
	}
	for i := 0; i < 10000; i++ {
		i := i
		go func() {
			cache.AddBalance(strconv.Itoa(i), float64(i)*200.)
			fmt.Printf("Adding %v with value %v\n", strconv.Itoa(i), float64(i)*200.)
		}()
	}
	go func() {
		averageBalance := cache.AverageBalance()
		fmt.Println("Average balance: ", averageBalance)
	}()

	go func() {
		fmt.Println("Map Size: ", cache.Size())
	}()
}
