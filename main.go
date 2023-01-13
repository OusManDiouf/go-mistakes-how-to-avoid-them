// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	fmt.Printf("s1 %v, len(%d), cap(%d)\n", s1, len(s1), cap(s1))
	s2 := s1[1:2]
	fmt.Printf("s2 %v, len(%d), cap(%d)\n", s2, len(s2), cap(s2))
	s3 := append(s2, 10)
	fmt.Printf("s3 %v, len(%d), cap(%d)\n", s3, len(s3), cap(s3))
	fmt.Printf("s1 %v, len(%d), cap(%d)\n", s1, len(s1), cap(s1))
	fmt.Printf("s1 %p\n", s1)
	fmt.Printf("s2 %p\n", s2)
	fmt.Printf("s3 %p\n", s3)
	fmt.Println("----------------------------")
	ss1 := []int{1, 2, 3}
	fmt.Printf("ss1 %v, len(%v), cap(%v)\n", ss1, len(ss1), cap(ss1))
	fmt.Printf("ss1 %p\n", ss1)
	// Full Slice Expression limit la capacite a 2,
	// donc un append ulterieur va creer un nouveau backing array !
	// [low, high, high-low]
	ss2 := ss1[:2:2]
	fmt.Printf("ss2 %p\n", ss2)
	f(ss2)
	fmt.Printf("ss1 %v, len(%v), cap(%v)\n", ss1, len(ss1), cap(ss1))
}
func f(ss2 []int) {
	fmt.Printf("ss2 %v, len(%v), cap(%v)\n", ss2, len(ss2), cap(ss2))
	ss2 = append(ss2, 100)
	fmt.Printf("ss2 %p\n", ss2)
	fmt.Printf("ss2 %v, len(%v), cap(%v)\n", ss2, len(ss2), cap(ss2))

}
