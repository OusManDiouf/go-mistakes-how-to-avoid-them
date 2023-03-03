package main

import (
	"fmt"
	"sync"
)

func main() {

	//------------------------ GO MEMORY MODEL ----------------------

	GoMemModelGuarantee5()
}

func GoMemModelGuarantee5() {
	// Unbuffered channel :
	//	La goroutine émeteur bloque jusqu'à ce que la goroutine du récepteur soit prête.
	// Buffered channel :
	//	La goroutine de l'émetteur ne se bloque que lorsque le tampon est plein.

	// Cette guarantie pourrait se montrer contre-intuitive, car elle statue que :
	// une reception à partir d'un unbuffered channel a lieu AVANT
	// que l'envoie vers ce channel se complete !

	// OPTION I : With buffered channel
	//	 La goRoutine parent envoie un message puis lit la variable i,
	//	 pendant que la goRoutine fille modifie la variable i et lit le message du parent
	//	 ici on remarque qu'il y a data race, car le read/write peut arriver simultanement,
	//	 la variable i n'est synchronisé.
	// OPTION II : With unbuffered channel
	//	 Dans ce cas de figure, le write est garantie de se produire avant le read,
	//	 parce que la reception à partir d'un channel unbufferisé ce produit AVANT
	//	 l'envoie vers ce même channel, donc le write sur i se fera toujour
	//	 avant le read de ce dernier dans la goRoutine parent.
	//	 Cet ordre est garanti par le "Go Memory Model"

	i := 0
	// first:  sender block until receive is ready
	// second: reveive happen before send complete
	ch := make(chan struct{})

	//ch := make(chan struct{}, 1)
	go func() {
		// [with Unbuffered Channel] this "write"
		// is guaranteed to happen before the read on the ch
		i = 1
		fmt.Println(<-ch) // receive from ch
	}()

	// this will block with an unbuff chan
	// but proceed with a buff chan (data race here)
	ch <- struct{}{}
	fmt.Println("[DONE] i = ", i)
}
func GoMemModelGuarantee4() {
	// La fermeture d'un channel a lieu avant reception de la notif de fermeture.
	// Le channel sera fermé une fois que le receiver receptionne
	// la derniére valeur retournée par ce channel.
	// Ici, au lieu d'envoyer un message, on ferme le channel
	// Ce cas de figure est exempt de data race.
	i := 0
	isChClosed := func(t bool) string {
		if t == false {
			return "YES"
		}
		return "NO"
	}
	ch := make(chan struct{})
	go func() {
		_, ok := <-ch
		fmt.Println("is ch closed: ", isChClosed(ok))
		fmt.Println("i: ", i)
	}()
	i++
	// A call to a close(ch) should always happen on the sender,
	// never on the receiver.
	close(ch)
}
func GoMemModelGuarantee3() {
	// SEND TO A CHANNEL HAPPEN BEFORE THE RECEIVE FROM THAT CHANNEL COPLETE.

	// Dans l'exemple suivant, une G parent incrémente une variable avant un envoi,
	// tandis qu'une autre G le lit après une lecture sur ce canal :

	// la goRoutine parent inc i avant l'emission vers un channel,
	// tandis qu'un autre goRoutine lit la variable i juste aprés un channel read

	// orderd'exec: variable increment < channel send < channel receive < varible read
	// par transitivité, on peut s'assurer que l'accés à i est synchronisé,
	// du coup pas de data race
	i := 0
	wg := sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan struct{}) // Unbuff chan - G sender will block untill the G receiver is ready.
	go func() {               //G1
		fmt.Println("[PRINT FIRST]ch res: ", <-ch) // receive on ch channnel
		fmt.Println("[PRINT SECOND]i: ", i)
		defer wg.Done()
	}()
	i++
	// here the main G will block until the above G receiver is ready
	// Also the "send" will happen before any receive on ch channel
	ch <- struct{}{}

	// First : any write in G1 will force this to be synchronized via wg to avoid data race
	// Second : l'affichage est random, ce qui veut dire que ce Print ne s'affiche pas en dernier
	// tout le temps, pour le rendre deterministic, on doit introduire un wg
	// pour avoir une guarantie d'afficher cette ligne tout le temps en dernier lieu
	wg.Wait()
	fmt.Println("------- [PRINT LAST] --------")
}
func GoMemModelGuarantee2() {
	// Cependant, la sortie d'une goRoutine n'est pas guarantie
	// de se dérouler avant un qulquonque event.
	// C'est pourquoi l'exemple suivant provoque un "data race".
	// Si l'on veut empêcher cela, on doit synchroniser les goRoutines
	i := 0
	go func() {
		i++
	}()
	fmt.Println("i: ", i) // syscall - data race - need sync via wg
}

func GoMemModelGuarantee1() {
	// Creating a goroutine happen before the goroutine exec begins.
	// Data race free on this one; the read happen before the G execution.
	// Donc, le fait de lire une variable et d'ensuite créer une goRoutine
	// qui ecrase cette variable ne provoque pas de data race.
	i := 0
	go func() {
		i++
		fmt.Println("i: ", i)
	}()
}
