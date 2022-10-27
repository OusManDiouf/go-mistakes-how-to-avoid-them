package main

import (
	"fmt"
)

func main() {
	guarantee5()
}

func guarantee5() {
	// unbuffered channel: La goroutine de l'émetteur bloque jusqu'à ce que la goroutine du récepteur soit prête.
	// buffered channel  : La goroutine de l'émetteur ne se bloque que lorsque le tampon est plein.

	// Cette dernière guarantie pourrait se montrer contre-intuitive.
	// une reception à partir d'un channel unbuffered à lieu AVANT
	// que l'envoie vers ce channel se complete!

	// OPTION I: with buffered channel
	// la goRoutine parent envoie un message pis lit la variable i,
	// pendant que la goRoutine fille modifie la variable i et lit le message du parent
	// ici on remarque qu'il y a data race car le read/write peut arrivé simultanement,
	// la variable i n'est synchronisé.
	// OPTION I: with unbuffered channel
	// dans ce cas, le write est garantie de se produire avant le read.
	// car la reception à partir d'un channel unbufferisé se produit AVANT
	// l'envoie vers ce même channel, donc le write sur i se fera toujour
	// avant le read de ce dernier dans la goRoutine parent.
	// Cette ordre est garantie par le "Go Memory Model"
	i := 0
	//ch := make(chan struct{}, 1)
	ch := make(chan struct{})
	go func() {
		i = 1
		fmt.Println(<-ch)
	}()
	ch <- struct{}{}
	fmt.Println(i)
}
func guarantee4() {
	// la fermeture d'un channel à lieu avant reception de la notif de fermeture.
	// le channel sera fermé une fois que le receiver receptionne la derniére value
	// retourner via ce channel.
	// ici, au lieu d'envoyer un message on ferme le channel
	// Ce cas de figure est exempt de data race.
	i := 0
	ch := make(chan struct{})
	go func() {
		_, ok := <-ch
		fmt.Println("ok to go on ch: ", ok)
		fmt.Println("i: ", i)
	}()
	i++
	close(ch)
}
func guarantee3() {
	// L'envoie vers un channel à lieu avant que le "receive from that channel" soit "complete"
	// la goRoutine parent inc i avant l'envoie ver un channel, tandis qu'un autre goRoutine,
	// lit la variable i juste aprés un channel read
	// orderd'exec: variable increment < channel send < channel receive < varible read
	// par transitivité, on peut s'assurer que l'accés à i est synchronisé,
	// du coup pas de data race
	i := 0
	ch := make(chan struct{})
	go func() {
		fmt.Println("ch res: ", <-ch)
		fmt.Println("i: ", i)
	}()
	i++
	ch <- struct{}{}
}
func guarantee2() {
	// Cependant l'exit d'une goRoutine n'est pas guarantie de se dérouler avant un event.
	// C'est pourquoi l'exemple suivant provoque un data race.
	// Si l'on veut empêcher cela, on doit synchroniser les goRoutines
	i := 0
	go func() {
		i++
	}()
	fmt.Println("i: ", i)
}

func guarantee1() {
	// la creation d'une goRoutine se passe avant le debut d'execution de la goRoutine.
	// Donc, le fait de lire une variable et d'ensuite créer une goRoutine
	// qui ecrase cette variable ne provoque pas de data race.
	i := 0
	go func() {
		i++
		fmt.Println("i: ", i)
	}()
}
