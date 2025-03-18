package main

import (
	"fmt"
)

const (
	end        = 10000
	numThreads = 4
)

var (
	counter = 0
)

func main() {
	array := [end]int{} // Ο πίνακας τώρα είναι τοπικά στη main και περναέι σαν όρισμα σε κάθε goroutine

	for i := 0; i < numThreads; i++ {
		go func(id int, array *[end]int) {
			for {
				if counter >= end {
					break
				}
				array[counter]++
				counter++
			}
			fmt.Printf("Thread %d finished with counter = %d\n", id, counter)
		}(i, &array)
	}

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
