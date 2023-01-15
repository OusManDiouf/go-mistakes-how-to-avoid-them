/**
Just like with slices, if we know up front the number of elements a map will contain,
we should create it by providing an initial size. Doing this avoids potential map growth,
which is quite heavy computation-wise because it requires reallocating enough space and rebalancing all the elements.
*/
package main

import "fmt"

func main() {
	fmt.Println("Go!!!!!")
}
