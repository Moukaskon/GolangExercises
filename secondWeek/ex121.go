package main

import (
	"fmt"
	"sync"
)

const (
	end = 1000
	numThreads = 4
)

// SharedData struct to hold the shared array
type SharedData struct {
	array [end]int
}

func main() {
	var data SharedData
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(id int, shared *SharedData) {
			defer wg.Done()
			for i := 0; i < end; i++ {
				for j := 0; j < i; j++ {
					shared.array[i]++
				}
			}
			fmt.Printf("Thread %d finished.\n", id)
		}(i, &data)
	}

	wg.Wait()

	checkArray(data)
}

func checkArray(data SharedData) {
	errors := 0
	fmt.Println("Checking...")

	for i := 0; i < end; i++ {
		if data.array[i] != numThreads*i {
			errors++
			fmt.Printf("%d: %d should be %d\n", i, data.array[i], numThreads*i)
		}
	}
	fmt.Println(errors, "errors.")
}
