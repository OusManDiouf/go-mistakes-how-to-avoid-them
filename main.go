package main

import "fmt"

func main() {

	m := map[int]bool{
		1: true,
		2: false,
		3: true,
	}
	fmt.Println(m)
	wrongInsertOnInterationMethod(m)
	fmt.Println(m)
}

func wrongInsertOnInterationMethod(m map[int]bool) {
	for k, v := range m {
		if v {
			m[10+k] = true
		}
	}
}
func goodInsertOnInterationMethod(m map[int]bool) {

}
