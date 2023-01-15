/**
If performance is paramount, implementing a custom method for comparaison is the way to go.
It's only a valid case, where we are dealing with types that are not comparable using ==
as it is for slice and map.
*/
package main

import (
	"fmt"
	"math"
)

type customer struct {
	ID         string
	Operations []float32
}

func (c customer) equal(c1 customer) any {
	if c.ID != c1.ID {
		return false
	}
	if len(c.Operations) != len(c1.Operations) {
		return false
	}
	for i := 0; i < len(c.Operations); i++ {
		if c.Operations[i] != c1.Operations[i] {
			return false
		}
	}
	return true
}

func main() {
	cust1 := customer{ID: "cust1", Operations: []float32{1.}}
	sameCust1 := customer{ID: "cust1", Operations: []float32{1.}}
	//fmt.Println("Are the customer the same ? ", cust1 == sameCust1) // the compiler alert us !
	fmt.Println("Are the customer the same ? ", cust1.equal(sameCust1)) // the compiler alert us !
	fmt.Println(1.0001 * 1.0001)
	fmt.Println(math.SmallestNonzeroFloat32)
	fmt.Println(math.MaxFloat64)

}
