package main

import (
	"fmt"
	"sync"
)

const (
	end        = 1000
	numThreads = 4
)

var (
	array [end]int
	mu    sync.Mutex
)

func main() {
	threads := make([]*CounterThread, numThreads)

	var wg sync.WaitGroup
	wg.Add(numThreads)

	for i := 0; i < numThreads; i++ {
		threads[i] = &CounterThread{}
		go func(id int) {
			defer wg.Done()
			threads[id].Run()
		}(i)
	}

	wg.Wait()
	checkArray()
}

func checkArray() {
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

type CounterThread struct{}

func (c *CounterThread) Run() {
	for i := 0; i < end; i++ {
		for j := 0; j < i; j++ {
			mu.Lock()
			array[i]++
			mu.Unlock()
		}
	}
}
