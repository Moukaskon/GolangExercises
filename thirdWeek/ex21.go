// Επέλεξα τον κώδικα από την άσκηση ex13

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
	array [end]int32
	mutex [end]sync.Mutex
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < end; i++ {
				for j := 0; j < i; j++ {
					mutex[i].Lock()
					array[i]++
					mutex[i].Unlock()
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
		mutex[i].Lock()
		actual := array[i]
		mutex[i].Unlock()

		if actual != expected {
			errors++
			fmt.Printf("%d: %d should be %d\n", i, actual, expected)
		}
	}
	fmt.Println(errors, "errors.")
}
