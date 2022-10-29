package main

import (
	"fmt"
	"time"
)

func main() {
	//execPrinter()
	timingOutGoRoutine()
}
func timingOutGoRoutine() {
	c1 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "c1 OK"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(5 * time.Second):
		fmt.Println("TimeOut After 5s...Exit")
		return
	}
}

func execPrinter() {

	var ch = make(chan bool)

	// write 5 times to the channel and this one is closed
	go printer(ch, 5)

	// le range loop s'arrếte de lui même puis que le channel est férmé dans printer
	for val := range ch {
		fmt.Print(val, " ")
	}
	fmt.Println()

	// le channel est deja clos donc il renvoie juste la valeur par default de son type!
	for i := 0; i < 15; i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println()
}

// Une seule goRoutine qui écrit sur une channel puis ferme ce dernier
// ne peut pas provoquer de data race
func printer(ch chan<- bool, nTimes int) {
	for i := 0; i < nTimes; i++ {
		ch <- true
	}
	close(ch)
}
