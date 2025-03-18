package main

import (
	"fmt"
)

const (
	end = 1000
	numThreads = 4
)

var array [end]int

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
		if array[i] != numThreads*i {
			errors++
			fmt.Printf("%d: %d should be %d\n", i, array[i], numThreads*i)
		}
	}
	fmt.Println(errors, "errors.")
}

type CounterThread struct{}

func (c *CounterThread) Start() {
	go c.Run()
}

func (c *CounterThread) Run() {
	for i := 0; i < end; i++ {
		for j := 0; j < i; j++ {
			array[i]++
		}
	}
}
