package main

import (
	"fmt"
	"strconv"
	"sync"
)

type CacheV1 struct {
	//TODO: try with mu sync.Map
	mu       sync.RWMutex
	balances map[string]float64
}

func (c *CacheV1) AddBalance(id string, balance float64) {
	c.mu.Lock()
	c.balances[id] = balance
	c.mu.Unlock()
}

func (c *CacheV1) AverageBalance() float64 {
	c.mu.Lock()
	balances := c.balances
	// soluce: if the iteration operation isn't heavy(that's the case here,
	//we only perform an increment operation), we should protect the whole function
	defer c.mu.Unlock()

	sum := 0.
	for _, balance := range balances {
		sum += balance
	}
	return sum / float64(len(balances))
}

func (c *CacheV1) Size() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.balances)
}

func main() {
	cache := CacheV1{
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
