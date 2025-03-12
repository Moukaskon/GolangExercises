package main

// Η main υπάρχει σε πάνω από ένα αρχεία και έτσι το IDE έχει συντακτικό. Για να τρέξει σωστά θα πρέπει να
// αντιμετωπίζονται σαν ξεχωριστά files και όχι ένα project. Τα τρέχω ένα ένα στο GoPlayground.

// 1) Βάζοντας σε σχόλια τα κομμάτια κώδικα όπου χρησιμοποιήται το WaitGroup για να εκτελεστούν τα threads, δεν παρατηρείτε καμία διαφορά στην εκτέλεση του προγράμματος, 
// αν όμως αυξήσουμε των αριθμό των threads από 20 σε 50 τότε βλέπουμε πως σχεδόν ποτέ δεν τυπώνεται η φράση "Hello from goroutine...".
// Αυτό συμβαίνει γιατί το πρόγραμμα δεν περιμένει να ολοκληρωθεί το thread ώστε να τυπώσει το αποτέλεσμα. Επίσης τα threads τρέχουν στο background όσο τρέχει η main 
// οπότε χωρίς το wait group το πρόγραμμα τερματίζεται μόλις ολοκληρωθεί η main. Λογικά κάτι παρόμοιο θα συμβαίνει και στην java αν δεν χρησιμοποιήσουμε κάποια δομή δεδομένων.

import (
	"fmt"
)

func main() {
	numThreads := 20
	//var wg sync.WaitGroup

	//wg.Add(numThreads)

	for i := 0; i < numThreads; i++ {
		fmt.Printf("Now starting goroutine %d\n", i)
		go func(id int) {
			//defer wg.Done()
			fmt.Printf("Hello from goroutine %d\n", id)
		}(i)
	}

	//wg.Wait()
	fmt.Println("In main: goroutines all done")
}