package main

import (
	"fmt"
	"strconv"
	"sync"
)

type CacheV3 struct {
	mu       sync.Mutex
	balances sync.Map
	size     int
}

func (c *CacheV3) AddBalance(id string, balance float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.balances.Store(id, balance)
	c.size++
}

func (c *CacheV3) AverageBalance() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	sum := 0.
	c.balances.Range(func(key, value any) bool {
		sum += value.(float64)
		return true
	})
	return sum / float64(c.size)
}

func (c *CacheV3) Size() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.size
}

func main() {
	cache := CacheV3{
		balances: sync.Map{},
		size:     0,
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
