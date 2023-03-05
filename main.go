package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// An empty struct is a de facto standard to convey an absence of meaning
	var s struct{}
	var i interface{}
	fmt.Println("SizeOf struct{} : ", unsafe.Sizeof(s))
	fmt.Println("SizeOf interface{} : ", unsafe.Sizeof(i))

	// Applied to channels, if we want to create a channel to send notifications without data,
	// the appropriate way to do so in Go is a chan struct{}.

	// If we want to design an idiomatic API in regard to Go standards,
	// let’s remember that a channel without data should be expressed with a chan struct{} type.
}
