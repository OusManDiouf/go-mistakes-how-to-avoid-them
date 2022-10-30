package main

func main() {
	/**
		Lorsqu'on utilise append le comportement depent du fait que le slice
	    soit plein (len == cap). Si c'est le cas, le runtime Go crée une nouveau
	  	backing array pour ajouter le nouvel element;
		autrement,le runtime l'ajoute au backing array existant.
		dans notre exemple, les deux go routine utilise le slice deja rempli à ras bord,
		donc dans les deux cas, nouveau backing array est créé dans chaque goRoutine,
		de ce fait, l'array existant n'est pas muté
		Conclusion: Pas de data race.
	*/
	//####################################################
	//sliceWithoutDataRace := make([]int, 1)
	//
	//go func() {
	//	s1 := append(sliceWithoutDataRace, 1)
	//	fmt.Println(s1)
	//}()
	//go func() {
	//	s2 := append(sliceWithoutDataRace, 1)
	//	fmt.Println(s2)
	//}()
	//go func() {
	//	fmt.Println(sliceWithoutDataRace)
	//}()
	//####################################################
	/**
		Dans le deuxieme exemple on utilise une slice avec un len=0 et un cap de 1
		dans ce cas, les deux goRoutine councours à MaJ le seule index dispo du tableau.
		autrement dit, les deux goRoutines essaye d'accéder à la même adresse memoire
		Conclusion: data race imminent !
		Solution:
			-> crée une copie de s dans chaque goRoutine, puis operer dessus
	           au lieu de toucher au slice original.
	*/
	//sliceWithDataRace := make([]int, 0, 1)
	//go func() {
	//	s1 := append(sliceWithDataRace, 1)
	//	fmt.Println(s1)
	//}()
	//go func() {
	//	s2 := append(sliceWithDataRace, 1)
	//	fmt.Println(s2)
	//}()
	//####################################################
	/**
	Accessing different slice indices regardless of the operation isn’t a data race;
	different indices mean different memory locations.
	*/
	sliceWithoutDataRaceDiffAdressIndexes := make([]int, 0, 1)
	sliceWithoutDataRaceDiffAdressIndexes = append(sliceWithoutDataRaceDiffAdressIndexes, 0)
	sliceWithoutDataRaceDiffAdressIndexes = append(sliceWithoutDataRaceDiffAdressIndexes, 0)
	go func() {
		sliceWithoutDataRaceDiffAdressIndexes[0] = 1
	}()
	go func() {
		sliceWithoutDataRaceDiffAdressIndexes[1] = 1
	}()

}
