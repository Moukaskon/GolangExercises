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

var array [end]int32

func main() {
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < end; i++ {
				for j := 0; j < i; j++ {
					atomic.AddInt32(&array[i], 1)
				}
			}
			fmt.Printf("Thread %d finished.\n", id)
		}(i)
	}

	wg.Wait()
	checkArray()
}

func checkArray() {
	errors := 0
	fmt.Println("Checking...")

	for i := 0; i < end; i++ {
		expected := int32(numThreads * i)
		if atomic.LoadInt32(&array[i]) != expected {
			errors++
			fmt.Printf("%d: %d should be %d\n", i, array[i], expected)
		}
	}
	fmt.Println(errors, "errors.")
}
