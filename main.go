package main

import "fmt"

func main() {

	theArray := []int{14, 25, 35, 47, 11, 25, 82, 5, 1, 3}
	fmt.Println(theArray)
	mergeSort(theArray)
	fmt.Println(theArray)

}
