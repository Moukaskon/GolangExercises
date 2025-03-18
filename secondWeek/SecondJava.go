package main

import (
	"fmt"
)

const (
	end = 10000
	numThreads = 4
)

var (
	counter int
	array   [end]int
)

func main() {
	threads := make([]*CounterThread, numThreads)

	for i := 0; i < numThreads; i++ {
		threads[i] = &CounterThread{}
		threads[i].Start()
	}

	checkArray()
}

func checkArray() {
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

type CounterThread struct{}

func (c *CounterThread) Start() {
	go c.Run()
}

func (c *CounterThread) Run() {
	for {
		if counter >= end {
			break
		}
		array[counter]++
		counter++
	}
}