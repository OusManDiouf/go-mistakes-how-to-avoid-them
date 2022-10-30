package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	//TODO: try with mu sync.Map
	mu       sync.RWMutex
	balances map[string]float64
}

func (c *Cache) AddBalance(id string, balance float64) {
	c.mu.Lock()
	c.balances[id] = balance
	c.mu.Unlock()
}

func (c *Cache) AverageBalance() float64 {
	c.mu.Lock()
	balances := c.balances // wrong way of copying: shared reference
	c.mu.Unlock()

	sum := 0.
	for _, balance := range balances {
		sum += balance
	}
	return sum / float64(len(balances))
}

func main() {
	cache := Cache{
		mu:       sync.RWMutex{},
		balances: map[string]float64{},
	}
	go func() {
		cache.AddBalance("1001", 200.)
		cache.AddBalance("1002", 100.)
		cache.AddBalance("1003", 200.)
		cache.AddBalance("1003", 300.)
	}()
	go func() {
		averageBalance := cache.AverageBalance()
		fmt.Println(averageBalance)
	}()
}
