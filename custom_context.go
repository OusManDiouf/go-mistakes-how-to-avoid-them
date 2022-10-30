package main

import (
	"context"
	"time"
)

/**
 	Suivant l'interface Context Originel
	-> le deadline du context est géré par la méthode Deadline
	-> le cancellation est géré via les méthodes Done et Err.
	-> Lorsque la deadline est expiré ou que le context est annulé,
	   Done devrait retourner un channel clos, et Err , une erreur.
	-> les values sont géré via la méthode Value

	Notez que cette example est inspiré de l'exemple context.go
	un emptyCtx est utilisé mais non exposé...

	On peut utiliser un custom context comme celui qui suit,
	ou bien utiliser un background context et lui assigné les Values
	du context parent, comme le fait l'implementation dans context.go
	( Background et Todo) sont des emptyCtx !
*/
//detach An emptyCtx is never canceled, has no values, and has no deadline.
//		 Custom context qui detache le signal d'annulation d'un context parent
// 		 ici detach est un ctx qui agit comme un wrapper autour d'un ctx parent.
type detach struct {
	ctx context.Context
}

// Deadline signal d'annulation detaché, donc return des valeur defaut
func (d detach) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

// Done signal d'annulation detaché, donc return nil
func (d detach) Done() <-chan struct{} {
	return nil
}

// Err signal d'annulation detaché, donc return nil
func (d detach) Err() error {
	return nil
}

// Value processing délégué au parent, puisqu'on veut garder les valeurs
func (d detach) Value(key any) any {
	return d.ctx.Value(key)
}
