package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
)

// Usefull stuf from aws go client sdk
const (
	// Byte is 8 bits
	Byte int64 = 1
	// KiloByte (KiB) is 1024 Bytes
	KiloByte = Byte * 1024
	// MebiByte (MiB) is 1024 KiB
	MebiByte = KiloByte * 1024
	// GibiByte (GiB) is 1024 MiB
	GibiByte = MebiByte * 1024
)

func main() {
	file, err := os.Open("100MB.bin")
	if err != nil {
		log.Fatalf("File can't be opened: %v\n", err)
	}

	//chunkCount, err := ReadSequential(file, KiloByte)
	//if err != nil {
	//	log.Fatalf("Error while reading file: %v\n", err)
	//	return
	//}
	//fmt.Println("Number of chunk Read with pooled version : ", chunkCount)

	chunkCount, err := ReadPooled(file, KiloByte)
	if err != nil {
		log.Fatalf("Error while reading file: %v\n", err)
		return
	}
	fmt.Println("Number of chunk Read with pooled version : ", chunkCount)
}

func ReadPooled(r io.Reader, buffSize int64) (int, error) {

	var chunkCount int64
	wg := sync.WaitGroup{}
	var poolSize = runtime.GOMAXPROCS(0)

	ch := make(chan []byte, poolSize)
	wg.Add(poolSize)
	// spin up n goRoutines (workers)
	for i := 0; i < poolSize; i++ {
		GoroutineID := i + 1
		go func() {
			wg.Done()
			for b := range ch {
				task(GoroutineID, b)
				atomic.AddInt64(&chunkCount, 1)
			}
		}()
	}

	// loop over the reader content and pull 1024 bytes on each step
	// and send it to workers via the shared channel ch
	for {
		b := make([]byte, buffSize)
		// read from r
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		// publish to the channel after every read
		ch <- b
	}
	close(ch)
	wg.Wait()

	return int(atomic.LoadInt64(&chunkCount)), nil
}

func ReadSequential(r io.Reader, buffSize int64) (int, error) {
	chunkCount := 0
	for {
		buff := make([]byte, buffSize)
		_, err := r.Read(buff)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		task(0, buff)
		chunkCount++
	}

	return chunkCount, nil
}

func task(workerNumber int, b []byte) {
}
