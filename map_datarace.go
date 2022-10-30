package main

func main() {

	/*
		Le fait d'acceder à la même Map (que ce soit avec la même ou une clé differente)
		via au moins deux goRoutine avec une goRoutine qui MaJ la Map represent un data race
		Pourquoi ? il faut savoir qu'une map est representé par un tableau de sceaux (array of buckets)
		dont chaque sceau (bucket) est un pointer vers un tableau de paires de key/value.
		un algo de hashing est utilisé pour determiner l'index d'un bucket.
		Puisque l'algo contient un certain caractère aleatoire durant l'initialisation de la map,
		une execution pourrait mener ou non à la même index.
		le detecteur de data race gére ce cas de figure en remontant un warning,
		independament du fait que le data race à bien eut lieu ou pas.
	*/

	m := make(map[int]bool, 2)
	go func() {
		m[0] = true
	}()
	go func() {
		m[1] = true
	}()
}
