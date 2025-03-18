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
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < end; i++ {
				for j := 0; j < i; j++ {
					mu.Lock()
					array[i]++
					mu.Unlock()
				}
			}
			fmt.Printf("Thread %d finished.\n", id)
		}(i)
	}

	wg.Wait()
	checkArray(array)
}

func checkArray(array [end]int) {
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
