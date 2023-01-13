package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

/**
☑ var s []string if we aren’t sure about the final length and the slice can be empty
☑ []string(nil) as syntactic sugar to create a nil and empty slice
☑ make([]string, length) if the future length is known
*/

type customer struct {
	ID         string
	Operations []float32
}

func main() {
	// EMPTY & NIL SLICE (NO ALLOCATION)
	var s []string
	log(1, s)
	s = []string(nil)
	log(2, s)

	// EMPTY & NON NIL SLICE (WITH ALLOCATION)
	s = []string{}
	log(3, s)
	s = make([]string, 0)
	log(4, s)

	log(5, f())
	log(6, intToString([]int{1, 2, 3, 4}))

	// Distinction between nil and empty slice
	var s1 []float32
	customer1 := customer{
		ID:         "foo",
		Operations: s1,
	}
	b1, _ := json.Marshal(customer1)
	fmt.Println(string(b1))

	s2 := make([]float32, 0)
	customer2 := customer{
		ID:         "bar",
		Operations: s2,
	}
	b2, _ := json.Marshal(customer2)
	fmt.Println(string(b2))
}

func log(i int, s []string) {
	fmt.Printf("%d : empty=%t\t nil=%t\n", i, len(s) == 0, s == nil)
}

func f() []string {
	isFoo := func() bool { return false }
	isBar := func() bool { return false }

	// To avoid unecessary alloc on the return in case none of the conditions are met,
	// we should initialize with an empty and nil slice.
	var s []string

	// this technique is useful for appending on a nil slice
	// s:= append([]string(nil), 42)

	// Do not allocate if the slice is potentialy nil!
	//s := make([]string, 0)

	// This technique is recommended to create a slice with initial elements
	//s := []string{"foo", "bar"}

	if isBar() {
		s = append(s, "bar")
	}
	if isFoo() {
		s = append(s, "foo")
	}

	return s
}

// In case we should produce a slice with a known length,
// we should use make([]string, anInt) to avoid extra allocation and copies.
func intToString(ints []int) []string {
	s := make([]string, len(ints))
	for i, v := range ints {
		s[i] = strconv.Itoa(v)
	}
	return s
}
