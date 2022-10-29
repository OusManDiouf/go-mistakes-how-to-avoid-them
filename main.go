package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatalf("File can't be opened: %v\n", err)
	}
	_, err = read(file)
	if err != nil {
		log.Fatalf("Error while reading file: %v\n", err)
		return
	}
}

func read(r io.Reader) (int, error) {

	var count int64
	wg := sync.WaitGroup{}
	var poolSize = runtime.GOMAXPROCS(0)

	ch := make(chan []byte, poolSize)
	wg.Add(poolSize)
	// spin up n goRoutines (workers)
	for i := 0; i < poolSize; i++ {
		loopCount := i + 1
		go func() {
			wg.Done()
			for b := range ch {
				_ = task(loopCount, b)
				//v := task(loopCount, b)
				//atomic.AddInt64(&count, int64(v))
			}
		}()
	}

	// loop over the reader content and pull 1024 bytes on each step
	// and send it to workers via the shared channel ch
	for {
		b := make([]byte, 1024)
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

	return int(count), nil
}

func task(workerNumber int, b []byte) int {
	fmt.Printf("Worker#%v process %v bytes\n", workerNumber, len(b))
	return 0
}
