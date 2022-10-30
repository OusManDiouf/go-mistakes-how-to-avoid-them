package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}

	for _, i := range s {
		//SOLUCE#1
		//val := i
		//go func() {
		//	fmt.Println(val)
		//}()
		//SOLUCE#2
		go func(val int) {
			fmt.Println(val)
		}(i)
	}
}
