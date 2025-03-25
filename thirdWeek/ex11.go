package main

// Αντικατέστησα το mutex με το atomic. Είναι πολύ πιο γρήγορο και απλό για μικρά tasks
// αλλά όχι για πολύπλοκα critical sections. Μπορεί να υποκαταστήσει το sync.Mutex.

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
	wg.Add(numThreads)

	for i := 0; i < numThreads; i++ {
		go func() {
			defer wg.Done()
			run()
		}()
	}

	wg.Wait()
	checkArray()
}

func run() {
	for i := 0; i < end; i++ {
		for j := 0; j < i; j++ {
			atomic.AddInt32(&array[i], 1)
		}
	}
}

func checkArray() {
	errors := 0
	fmt.Println("Checking...")

	for i := 0; i < end; i++ {
		expected := numThreads * i
		if int(array[i]) != expected {
			errors++
			fmt.Printf("%d: %d should be %d\n", i, array[i], expected)
		}
	}
	fmt.Println(errors, "errors.")
}
