package main

import (
	"fmt"
	"sync"
)

const (
	end        = 10000
	numThreads = 4
)

var (
	counter int
	array   [end]int
	mu      sync.Mutex
)

func main() {
	threads := make([]*CounterThread, numThreads)
	var wg sync.WaitGroup

	wg.Add(numThreads)

	for i := 0; i < numThreads; i++ {
		threads[i] = &CounterThread{}
		go func() {
			defer wg.Done()
			threads[i].Run()
		}()
	}

	wg.Wait()
	checkArray()
}

func checkArray() {
	errors := 0
	fmt.Println("Checking...")

	for i := 0; i < end; i++ {
		if array[i] != 1 {
			errors++
			fmt.Printf("%d: %d should be 1\n", i, array[i])
		}
	}
	fmt.Println(errors, "errors.")
}

type CounterThread struct{}

func (c *CounterThread) Run() {
	for {
		mu.Lock()
		if counter >= end {
			mu.Unlock()
			break
		}
		array[counter]++
		counter++
		mu.Unlock()
	}
}
