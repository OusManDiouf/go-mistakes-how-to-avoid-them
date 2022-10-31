package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go!!!!!")
}
func getCount() int {
	i := 0
	//wg := sync.WaitGroup{}
	//wg.Add(1)

	go func() {
		//defer wg.Done()
		i++
	}()
	//wg.Wait()
	return i
}
