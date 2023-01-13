package main

import (
	"fmt"
)

func main() {
	var s1 []float32
	customer1 := customer{
		ID:         "cust1",
		Operations: s1,
	}
	customer1.handleOperations()

	s2 := make([]float32, 0)
	customer2 := customer{
		ID:         "cust2",
		Operations: s2,
	}
	customer2.handleOperations()
}

type customer struct {
	ID         string
	Operations []float32
}

//func (c *customer) addOperations(ops []float32) {
//	c.Operations = ops
//}
func (c *customer) handleOperations() {
	operations := c.getOperations()

	// Logique corompu du fait de la manière dont on a initialisé Operations.
	// Operations est tantôt néant, tantôt empty.
	// Dans les deux cas, il n'y a pas d'operations à handle,
	// alors que l'un des cas invoque la fonction handle même s'il n'y a pas d'operations !
	//if operations != nil {
	//	c.handle()
	//	return
	//}

	// Donc si on veut vérifier qu'un slice est néant ou empty,
	// On utilise la fonction len(), dans les deux cas, aucun processing ne sera effectué.
	if len(operations) != 0 {
		c.handle()
	}

	fmt.Printf("%+v : No operations to handle for %s.\n", operations, c.ID)
}
func (c *customer) getOperations() []float32 {

	// On peut modifier le callee pour qu'il retourne un slice néant
	// dans le cas ou c.operations est néant ou empty peu importe
	// la facon dont on la init au depart.
	// Dans ce cas, on peut utiliser la premiere semantique
	// au niveau de l'appelant (voir handleOperations)
	//if len(c.Operations) == 0 {
	//	return nil
	//}

	return c.Operations
}

func (c *customer) handle() {
	fmt.Printf("%+v : All operations for %s get handled successfuly. \n", c.Operations, c.ID)
}
