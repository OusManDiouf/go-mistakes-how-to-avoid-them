package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func BenchmarkReadSequential(b *testing.B) {
	// FOR THOSE FILE LOOK AT MAIN THINGY DISK
	//file, err := os.Open("1GB.bin")
	//file, err := os.Open("10GB.bin")
	file, err := os.Open("100MB.bin")
	if err != nil {
		log.Fatalf("File can't be opened: %v\n", err)
	}

	chunkCount, err := ReadSequential(file, MebiByte)
	if err != nil {
		log.Fatalf("Error while reading file: %v\n", err)
		return
	}
	fmt.Println("[sequential version]  - chunknumber = : ", chunkCount)
}

func BenchmarkReadPooled(b *testing.B) {
	// FOR THOSE FILE LOOK AT MAIN THINGY DISK
	//file, err := os.Open("1GB.bin")
	//file, err := os.Open("10GB.bin")
	file, err := os.Open("100MB.bin")
	if err != nil {
		log.Fatalf("File can't be opened: %v\n", err)
	}

	chunkCount, err := ReadPooled(file, MebiByte)
	if err != nil {
		log.Fatalf("Error while reading file: %v\n", err)
		return
	}
	fmt.Println("[pooled version]  - chunknumber = : ", chunkCount)

}
