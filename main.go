package main

import "fmt"

type Customer struct {
	ID      string
	Balance float32
}
type Store struct {
	m map[string]*Customer
}

//func (s *Store) storeCustomersMap(mCustomers map[int]*Customer) {
func (s *Store) storeCustomersMap(mCustomers map[int]Customer) {
	//for _, customer := range mCustomers {
	//	fmt.Printf("customer loop variable: %p\n", &customer)
	//	s.m[customer.ID] = &customer
	//}
	for i := range mCustomers {
		currentCustomer := mCustomers[i]
		s.m[currentCustomer.ID] = &currentCustomer
	}
}

//func (s *Store) storeCustomersSlice(customers ...*Customer) {
func (s *Store) storeCustomersSlice(customers ...Customer) {
	// la boucle crée une seule variable customer avec une addresse fixe!
	// peut importe le nombre d'elements...
	// Et c'est cette variable qui est utilisé à chaque boucle
	for _, customer := range customers {
		fmt.Printf("customer loop variable: %p\n", &customer)
		s.m[customer.ID] = &customer
		/**
		first iter: customer = & cust1 and this address is stored in the map
		second iter: customer = & cust2 and this address is stored in the map
		The has created a single customer variable with a fixed address
		regardless of the number of elements
		At the end of the iteration,our map have stored the same pointer 3 times.
		This pointer's last assignment is a reference to the slice's last element
		THIS IS WHY ALL THE MAP ELEMENT REFERENCE THE SAME CUSTOMER !
		*/
	}
	//Soluce
	//for i := range customers {
	//	s.m[customers[i].ID] = &customers[i]
	//}
}

func (s *Store) viewStore() {
	for key, customer := range s.m {
		fmt.Printf("Key:%s, Value:%v, Addr: %p\n", key, customer, s.m[key])
	}
}

func main() {
	cust1 := Customer{ID: "alan", Balance: 100.}
	cust2 := Customer{ID: "dave", Balance: 200.}
	store := Store{m: make(map[string]*Customer)}

	//mCustomers := map[int]*Customer{1: &cust1, 2: &cust2}
	mCustomers := map[int]Customer{1: cust1, 2: cust2}

	//store.storeCustomersSlice(cust1, cust2)
	//store.storeCustomersSlice(&cust1, &cust2)
	store.storeCustomersMap(mCustomers)
	store.viewStore()

}
