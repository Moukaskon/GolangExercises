package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	end = 10000
	numThreads = 4
)

type SharedData struct {
	array   [end]int32
	counter int32
}

func main() {
	var data SharedData
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(id int, shared *SharedData) {
			defer wg.Done()
			for {
				// Atomically increment the counter and get the index
				idx := int(atomic.AddInt32(&shared.counter, 1)) - 1
				if idx >= end {
					break
				}
				// Atomically increment the array at the given index
				atomic.AddInt32(&shared.array[idx], 1)
			}
			fmt.Printf("Thread %d finished with counter = %d\n", id, atomic.LoadInt32(&shared.counter))
		}(i, &data)
	}

	wg.Wait()
	checkArray(&data)
}

func checkArray(data *SharedData) {
	errors := 0
	fmt.Println("Checking...")

	for i := 0; i < end; i++ {
		// Use atomic load to read the value from the array
		if atomic.LoadInt32(&data.array[i]) != 1 {
			errors++
			fmt.Printf("%d: %d should be 1\n", i, data.array[i])
		}
	}
	fmt.Println(errors, "errors.")
}
