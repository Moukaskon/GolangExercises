package main

import (
	"fmt"
	"sync"
)

const (
	end        = 10000
	numThreads = 4
)

type SharedData struct {
	array   [end]int
	counter int
	mu      sync.Mutex
}

func main() {
	var data SharedData
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(id int, shared *SharedData) {
			defer wg.Done()
			for {
				shared.mu.Lock()
				if shared.counter >= end {
					shared.mu.Unlock()
					break
				}
				shared.array[shared.counter]++
				shared.counter++
				shared.mu.Unlock()
			}
			fmt.Printf("Thread %d finished with counter = %d\n", id, shared.counter)
		}(i, &data)
	}

	wg.Wait()

	checkArray(data)
}

func checkArray(data SharedData) {
	errors := 0
	fmt.Println("Checking...")

	for i := 0; i < end; i++ {
		if data.array[i] != 1 {
			errors++
			fmt.Printf("%d: %d should be 1\n", i, data.array[i])
		}
	}
	fmt.Println(errors, "errors.")
}
