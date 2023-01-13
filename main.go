package main

import "fmt"

func main() {
	src := []int{1, 2, 3}
	//var dest []int
	dest := make([]int, len(src))

	// Copy returns the number of elements copied,
	// which will be the minimum of len(src) and len(dst).
	copy(dest, src)
	fmt.Println("len(dest): ", len(dest))
	fmt.Println("len(src): ", len(src))
	fmt.Println("src: ", src)
	fmt.Println("dest: ", dest)
}
