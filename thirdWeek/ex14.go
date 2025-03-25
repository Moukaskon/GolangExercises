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

var counter int32

func main() {
	array := make([]int32, end)
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(id int, array []int32) {
			defer wg.Done()
			for {
				idx := int(atomic.AddInt32(&counter, 1)) - 1 
				if idx >= end {
					break
				}
				atomic.AddInt32(&array[idx], 1)
			}
			fmt.Printf("Thread %d finished with counter = %d\n", id, atomic.LoadInt32(&counter))
		}(i, array)
	}

	wg.Wait()
	checkArray(array)
}

func checkArray(array []int32) {
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
