package main

import (
	"fmt"
	"runtime"
)

//var store = map[int][]byte{} // look at the next mistake to see why it's inefficient
var store = make(map[int][]byte, 1_000)

type Foo struct {
	v []byte
}

/**
 * Demontre que le slicing peut mener à une fuite de mêmoire
 * à cause de la capacité du slice.
 */
func main() {
	// SLICE CAP LEAK
	consumeMessages()
	defer func() {
		printAlloc()
	}()

	// SLICE POINTER LEAK
	//foos := make([]Foo, 1_000)
	//printAlloc()
	//for i := 0; i < len(foos); i++ {
	//	foos[i] = Foo{v: make([]byte, 1024*1024)}
	//}
	//printAlloc()
	//two := keepFirstTwoElementsOnly(foos)
	//runtime.GC()
	//printAlloc()
	//runtime.KeepAlive(two)
}

func keepFirstTwoElementsOnly(foos []Foo) []Foo {
	//return foos[:2]

	// TECHNIQUE: (OPTIMAL) ici le GC sait que les 998 elements de foos ne seront plus referencés
	// et peuvent maintenant être collectés
	f := make([]Foo, 2)
	copy(f, foos)
	return f

	// TECHNIQUE: mettre à nil les elements indesirable tout en gardant la cap sous-jacente intacte
	//for i := 2; i < 1_000; i++ {
	//	foos[i].v = nil
	//}
	//return foos[:2]

}

func consumeMessages() {
	for i := 0; i < 1000; i++ {
		msg := receiveMessage()
		fmt.Printf("msg len(%v), cap(%v)\n", len(msg), cap(msg))
		// do some with msg
		storeMessageType(i, getMessageType(msg))
	}
}

func receiveMessage() []byte {
	return make([]byte, 1024*1024)
}

func storeMessageType(i int, messageType []byte) {
	store[i] = messageType
	fmt.Printf("store[%v] len(%v), cap(%v)\n", i, len(store[i]), cap(store[i]))

}
func getMessageType(msg []byte) []byte {
	// TECHNIQUE: Directly slicing msg
	// the cap of m is the same as the initial slice (msg)
	// so it keep a ref on it !
	// use 984.272 MB of mem !!!!!
	// m := msg[:5]

	// TECHNIQUE: copying the slice (OPTIMAL)
	/// here we make a copy instead of directly slicing msg
	// use 4.234 MB of mem !!!!!
	m := make([]byte, 5)
	// Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).
	copy(m, msg)

	// TECHNIQUE FULL SLICE EXPRESSION
	// GC est incapable de collecter l'espace restant du backing array de  msg,
	// du coup on hijack cette memoire alors qu'on en utilise qu'une petite portion via msg[:5:5]
	// le slice expression limite juste le nombre d'element à copier dans m
	// mais le backing array est toujours dispo en memoire !!
	//m := msg[:5:5]

	fmt.Printf("m len(%v), cap(%v)\n", len(m), cap(m))
	return m
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
