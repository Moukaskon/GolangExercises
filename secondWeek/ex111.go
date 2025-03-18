package main

// Εδώ βλέπουμε πως τα αποτελέσματα είναι όλα λανθασμένα καθώς η main τελειώνει πριν τις υπορουτίνες
// Αν βάλουμε τα threads σε ένα waitGroup τότε θα δούμε το πρόβλημα που δημιουργείται εξαιτίας των lock.
// Παρακάτω βάζω και αυτόν τον κώδικα.

import (
	"fmt"
)

const (
	end = 1000
	numThreads = 4
)

func main() {
	var array [end]int

	// Αρχίζω τα threads χωρίς πίνακα
	for i := 0; i < numThreads; i++ {
		go func(id int) {
			for i := 0; i < end; i++ {
				for j := 0; j < i; j++ {
					array[i]++
				}
			}
			fmt.Printf("Thread %d finished.\n", id)
		}(i)
	}

	checkArray(array)
}

func checkArray(array [end]int) {
	errors := 0
	fmt.Println("Checking...")

	for i := 0; i < end; i++ {
		if array[i] != numThreads*i {
			errors++
			fmt.Printf("%d: %d should be %d\n", i, array[i], numThreads*i)
		}
	}
	fmt.Println(errors, "errors.")
}


package main

import (
	"fmt"
	"sync"
)

const (
	end        = 1000
	numThreads = 4
)

func main() {
	var array [end]int
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < end; i++ {
				for j := 0; j < i; j++ {
					array[i]++
				}
			}
			fmt.Printf("Thread %d finished.\n", id)
		}(i)
	}

	wg.Wait()

	checkArray(array)
}

func checkArray(array [end]int) {
	errors := 0
	fmt.Println("Checking...")

	for i := 0; i < end; i++ {
		if array[i] != numThreads*i {
			errors++
			fmt.Printf("%d: %d should be %d\n", i, array[i], numThreads*i)
		}
	}
	fmt.Println(errors, "errors.")
}
