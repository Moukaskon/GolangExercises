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
	mu      sync.Mutex
)

func main() {
	array := [end]int{}
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(id int, array *[end]int) {
			defer wg.Done()
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
			fmt.Printf("Thread %d finished with counter = %d\n", id, counter)
		}(i, &array)
	}

	wg.Wait()
	checkArray(array)
}

func checkArray(array [end]int) {
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
