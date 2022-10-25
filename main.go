package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	StatusSuccess  = "success"
	StatusErrorFoo = "error_foo"
	StatusErrorBar = "error_bar"
)

func main() {

	logger := LogCounterMap{}
	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		err := f(logger)
		if err != nil {
			log.Println(err)
		}
	}
	fmt.Println(logger)
}

func f(counterMap LogCounterMap) error {
	var status string
	defer func() {
		// here stutus reference a variable outside the closure body
		// hence, status is evaluated when the closure is executed
		// and not when call defer.
		counterMap.incrementCounter(status)
	}()

	if err := foo(); err != nil {
		status = StatusErrorFoo
		return err
	}
	if err := bar(); err != nil {
		status = StatusErrorBar
		return err
	}

	status = StatusSuccess
	return nil
}

func foo() error {
	if 1 == rand.Intn(5) {
		return errors.New("err in foo")
	}
	return nil
}
func bar() error {
	if 2 == rand.Intn(5) {
		return errors.New("err in bar")
	}
	return nil
}

type LogCounterMap map[string]int

func (m LogCounterMap) incrementCounter(status string) {
	m[status] += 1
}
