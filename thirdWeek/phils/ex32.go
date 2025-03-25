package main

import (
	"sync"
)

const (
	numPhils   = 5
	sleepTime  = 500 // Σε milliseconds
)

func main() {
	var wg sync.WaitGroup
	wg.Add(numPhils)

	philosophers := make([]*Philosopher, numPhils)
	forks := make([]*Fork, numPhils)

	for i := 0; i < numPhils; i++ {
		forks[i] = NewFork(i)
	}

	for i := 0; i < numPhils; i++ {
		philosophers[i] = NewPhilosopher(i, (i+1)%numPhils, sleepTime, forks[i], forks[(i+1)%numPhils])
		go philosophers[i].Run(&wg)
	}

	wg.Wait()
}
