// Τρέχοντας το πρόγραμμα δεν παρατήρησα κάποια σημαντική διαφορά στην ταχύτητα
// εκτέλεσης του προγράμματος με την προσθήκη του threshold. Φαίνεται μια μικρή
// βελτίωση στην ταχύτητα εκτέλεσης, αλλά δεν είναι σημαντική. Δεν ξέρω αν γίνεται κάποιο λάθος στον κώδικα,
// αλλά αν μεγαλώσω κι άλλο το threshold το προγραμμα δεν τρέχει καθόλου, μάλλον λόγω του μηχανιματός μου.

package main

import (
	"fmt"
	"sync"
	"time"
)

func computePiParallel(start, end, totalSteps int, wg *sync.WaitGroup, resultChan chan float64, threshold int) {
	defer wg.Done()

	if end-start <= threshold {
		step := 1.0 / float64(totalSteps)
		sum := 0.0
		for i := start; i < end; i++ {
			x := (float64(i) + 0.5) * step
			sum += 4.0 / (1.0 + x*x)
		}
		resultChan <- sum
	} else {
		mid := (start + end) / 2
		wg.Add(2)
		go computePiParallel(start, mid, totalSteps, wg, resultChan, threshold)
		go computePiParallel(mid, end, totalSteps, wg, resultChan, threshold)
	}
}

func computePiDivideAndConquer(n int, threshold int) float64 {
	var wg sync.WaitGroup
	resultChan := make(chan float64, 1024)

	wg.Add(1)
	go computePiParallel(0, n, n, &wg, resultChan, threshold)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	sum := 0.0
	for val := range resultChan {
		sum += val
	}
	return sum * (1.0 / float64(n))
}

func main() {
	threshold := 10000000
	n := 1000000000
	for i := 100000; i <= threshold; i += 100000 {
		start := time.Now()
		pi := computePiDivideAndConquer(n, i)
		elapsed := time.Since(start)
		fmt.Printf("Divide-and-Conquer π = %.15f\n", pi)
		fmt.Printf("Time: %s\n", elapsed)
	}

}
