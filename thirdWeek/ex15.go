package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	end = 1000
	numThreads = 4
)

type SharedData struct {
	array [end]int32 // Use int32 for atomic operations
}

func main() {
	var data SharedData
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(id int, shared *SharedData) { // Pass pointer to avoid copying
			defer wg.Done()
			for i := 0; i < end; i++ {
				for j := 0; j < i; j++ {
					atomic.AddInt32(&shared.array[i], 1) // Atomic increment
				}
			}
			fmt.Printf("Thread %d finished.\n", id)
		}(i, &data)
	}

	wg.Wait()
	checkArray(&data)
}

func checkArray(data *SharedData) {
	errors := 0
	fmt.Println("Checking...")

	for i := 0; i < end; i++ {
		expected := int32(numThreads * i)
		if atomic.LoadInt32(&data.array[i]) != expected { // Atomic read
			errors++
			fmt.Printf("%d: %d should be %d\n", i, data.array[i], expected)
		}
	}
	fmt.Println(errors, "errors.")
}
