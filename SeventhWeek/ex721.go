// Παρατήρησα πως το Dynamic load balancing είναι πιο αργό από το Static load balancing. Αυτό μου έκανε εντύπωση και το έψαηα λίγο.
// Βρήκα πως στην ουσία το Dynamic load balancing είναι πιο αργό γιατί οι workers περιμένουν να τους στείλουμε δουλειά.
// Αντίθετα στο Static load balancing οι workers δουλεύουν παράλληλα και δεν περιμένουν. Είναι αυτός όμως ο μόνος λόγος?

package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"time"
)

func markMultiples(start, end int, prime int, sieve []bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		if i%prime == 0 {
			sieve[i] = false
		}
	}
}

func sieveStatic(n int, numWorkers int) []bool {
	sieve := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		sieve[i] = true
	}

	sqrtN := int(math.Sqrt(float64(n)))
	for p := 2; p <= sqrtN; p++ {
		if sieve[p] {
			var wg sync.WaitGroup
			chunk := (n - p*p) / numWorkers
			if chunk == 0 {
				chunk = 1
			}
			for i := 0; i < numWorkers; i++ {
				start := p*p + i*chunk
				end := start + chunk - 1
				if end > n {
					end = n
				}
				wg.Add(1)
				go markMultiples(start, end, p, sieve, &wg)
			}
			wg.Wait()
		}
	}
	return sieve
}

func sieveCyclic(n int, numWorkers int) []bool {
	sieve := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		sieve[i] = true
	}

	sqrtN := int(math.Sqrt(float64(n)))
	for p := 2; p <= sqrtN; p++ {
		if sieve[p] {
			var wg sync.WaitGroup
			for i := 0; i < numWorkers; i++ {
				wg.Add(1)
				go func(offset int) {
					defer wg.Done()
					for j := p*p + offset*p; j <= n; j += p * numWorkers {
						sieve[j] = false
					}
				}(i)
			}
			wg.Wait()
		}
	}
	return sieve
}

func sieveDynamic(n int, numWorkers int) []bool {
	sieve := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		sieve[i] = true
	}

	sqrtN := int(math.Sqrt(float64(n)))
	var mu sync.Mutex

	for p := 2; p <= sqrtN; p++ {
		if sieve[p] {
			jobChan := make(chan int, 1000)
			var wg sync.WaitGroup

			// Start workers
			for i := 0; i < numWorkers; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for j := range jobChan {
						mu.Lock()
						sieve[j] = false
						mu.Unlock()
					}
				}()
			}

			// Send jobs in a separate goroutine
			go func(p int) {
				for j := p * p; j <= n; j += p {
					jobChan <- j
				}
				close(jobChan)
			}(p)

			wg.Wait()
		}
	}

	return sieve
}

func writeTimeToFile(filename string, duration time.Duration) error {
	f, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			f, err = os.Create(filename)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
	}
	defer f.Close()

	_, err = f.WriteString(duration.String() + "\n")
	if err != nil {
		return err
	}

	return nil
}

func countPrimes(sieve []bool) int {
	count := 0
	for i := 2; i < len(sieve); i++ {
		if sieve[i] {
			count++
		}
	}
	return count
}

func main() {
	for i := 10; i < 100; i += 10 {
		n := 1000000 * i
		numWorkers := 8

		fmt.Printf("=== Static Load Balancing ===\n")
		start := time.Now()
		sieve := sieveStatic(n, numWorkers)
		duration := time.Since(start)
		fmt.Printf("Primes: %d | Time: %s\n\n", countPrimes(sieve), duration)
		writeTimeToFile("result.txt", duration)

		fmt.Printf("=== Cyclic Load Balancing ===\n")
		start = time.Now()
		sieve = sieveCyclic(n, numWorkers)
		duration = time.Since(start)
		fmt.Printf("Primes: %d | Time: %s\n\n", countPrimes(sieve), duration)
		writeTimeToFile("result.txt", duration)

		fmt.Printf("=== Dynamic Load Balancing ===\n")
		start = time.Now()
		sieve = sieveDynamic(n, numWorkers)
		duration = time.Since(start)
		fmt.Printf("Primes: %d | Time: %s\n", countPrimes(sieve), duration)
		writeTimeToFile("result.txt", duration)
	}
}
