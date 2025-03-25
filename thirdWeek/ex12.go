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

var (
	counter int32
	array   [end]int32
)

func main() {
	var wg sync.WaitGroup
	threads := make([]*CounterThread, numThreads)

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
		if atomic.LoadInt32(&array[i]) != 1 {
			errors++
			fmt.Printf("%d: %d should be 1\n", i, array[i])
		}
	}
	fmt.Println(errors, "errors.")
}

type CounterThread struct{}

func (c *CounterThread) Run() {
	for {
		idx := int(atomic.AddInt32(&counter, 1)) - 1
		if idx >= end {
			break
		}
		atomic.AddInt32(&array[idx], 1)
	}
}
