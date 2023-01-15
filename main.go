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

	//for _, a := range accounts {
	//	// the struct 'a' is a copy so its mutation is only local to the range scope!
	//	a.balance += 1000
	//}

	for i := range accounts {
		accounts[i].balance += 1000
	}

	// used when we only need to update some value
	//for i := 0; i < len(accounts); i++ {
	//	accounts[i].balance += 1000
	//}

	fmt.Println(accounts)

}
