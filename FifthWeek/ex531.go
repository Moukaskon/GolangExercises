// Και εδώ εφάρμοσα την ίδια αυτοματοποίηση. Θα πρέπει να ακολουθηθεί η διαδικασία
// της ex521 για να τρέξει ή να παέι σε νεο κενό φάκελο με δημιουργία mod.go file με αυτή
// την εντολή στο terminal: <go mod init ex531>.
// size 23 with 1 threads in 3690 ms
// size 23 with 2 threads in 1814 ms
// size 23 with 4 threads in 905 ms
// size 23 with 8 threads in 848 ms
// size 24 with 1 threads in 7355 ms
// size 24 with 2 threads in 3742 ms
// size 24 with 4 threads in 1949 ms
// size 24 with 8 threads in 1603 ms
// size 25 with 1 threads in 15068 ms
// size 25 with 2 threads in 7584 ms
// size 25 with 4 threads in 4192 ms
// size 25 with 8 threads in 3451 ms
// size 26 with 1 threads in 30322 ms
// size 26 with 2 threads in 15387 ms
// size 26 with 4 threads in 11062 ms
// size 26 with 8 threads in 27132 ms

package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {

	sizes := []int{23, 24, 25, 26}

	threadCounts := []int{1, 2, 4, 8}

	for _, size := range sizes {
		for _, threads := range threadCounts {
			processWithThreads(size, threads)
		}
	}
}

func processWithThreads(size int, numThreads int) {
	iterations := int(math.Pow(2, float64(size)))

	start := time.Now()

	var wg sync.WaitGroup

	chunkSize := iterations / numThreads
	if iterations%numThreads != 0 {
		chunkSize++
	}

	var output []string
	var mu sync.Mutex

	for t := 0; t < numThreads; t++ {
		startIdx := t * chunkSize
		endIdx := (t + 1) * chunkSize
		if endIdx > iterations {
			endIdx = iterations
		}

		wg.Add(1)

		go func(startIdx, endIdx int) {
			defer wg.Done()
			for i := startIdx; i < endIdx; i++ {
				checkCircuit(i, size, &output, &mu)
			}
		}(startIdx, endIdx)
	}

	wg.Wait()

	elapsed := time.Since(start)

	fmt.Printf("Processed size %d with %d threads in %d ms\n", size, numThreads, elapsed.Milliseconds())
}

func checkCircuit(z, size int, output *[]string, mu *sync.Mutex) {
	v := make([]bool, size)
	for i := size - 1; i >= 0; i-- {
		v[i] = (z & (1 << i)) != 0
	}

	value := (v[0] || v[1]) &&
		(!v[1] || !v[3]) &&
		(v[2] || v[3]) &&
		(!v[3] || !v[4]) &&
		(v[4] || !v[5]) &&
		(v[5] || !v[6]) &&
		(v[5] || v[6]) &&
		(v[6] || !v[15]) &&
		(v[7] || !v[8]) &&
		(!v[7] || !v[13]) &&
		(v[8] || v[9]) &&
		(v[8] || !v[9]) &&
		(!v[9] || !v[10]) &&
		(v[9] || v[11]) &&
		(v[10] || v[11]) &&
		(v[12] || v[13]) &&
		(v[13] || !v[14]) &&
		(v[14] || v[15]) &&
		(v[14] || v[16]) &&
		(v[17] || v[1]) &&
		(v[18] || !v[0]) &&
		(v[19] || v[1]) &&
		(v[19] || !v[18]) &&
		(!v[19] || !v[9]) &&
		(v[0] || v[17]) &&
		(!v[1] || v[20]) &&
		(!v[21] || v[20]) &&
		(!v[22] || v[20]) &&
		(!v[21] || !v[20]) &&
		(v[22] || !v[20])
	if value {
		saveResult(v, size, z, output, mu)
	}
}

func saveResult(v []bool, size, z int, output *[]string, mu *sync.Mutex) {
	result := fmt.Sprintf("%d", z)
	for i := 0; i < size; i++ {
		if v[i] {
			result += " 1"
		} else {
			result += " 0"
		}
	}

	mu.Lock()
	*output = append(*output, "\n"+result)
	mu.Unlock()

	//fmt.Println(result)
}
