// Δεν βρήκα κάποιο καλύτερο τρόπο να χωρίσω τα chunks. Για 1000000 στοιχεία κάνει 
// timeout στο playground γιαυτό το μείωσα σε 10000. Δεν έχει τεράστια διαφορά 
// από το μη παράλληλο της Java. Υπάρχει κάτι καλύτερο ή θα πρέπει να 
// μεγαλώσει πολύ το size για να δω διαφορά? (Μέχρι 500000 ήταν ελάχιστη)

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	size := 10000
	numWorkers := runtime.NumCPU()

	a := make([]float64, size)
	b := make([]float64, size)
	c := make([]float64, size)

	for i := 0; i < size; i++ {
		a[i] = 0.0
		b[i] = 1.0
		c[i] = 0.5
	}

	var wg sync.WaitGroup
	chunkSize := size / numWorkers

	for w := 0; w < numWorkers; w++ {
		start := w * chunkSize
		end := start + chunkSize
		if w == numWorkers-1 {
			end = size
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				a[i] = b[i] + c[i]
			}
		}(start, end)
	}

	wg.Wait()

	for i := 0; i < size; i++ {
		fmt.Println(a[i])
	}
}
