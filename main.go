package main

import "fmt"

type account struct {
	balance float32
}

func main() {
	accounts := []account{
		{balance: 100.},
		{balance: 200.},
		{balance: 300.},
	}

	for i := range accounts {
		accounts[i].balance += 1000
	}

	fmt.Println(accounts)

}
