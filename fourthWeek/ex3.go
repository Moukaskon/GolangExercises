package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Park struct {
	capacity int
	spaces   chan struct{}
}

func NewPark(capacity int) *Park {
	return &Park{
		capacity: capacity,
		spaces:   make(chan struct{}, capacity),
	}
}

func (p *Park) Arrive(id int) {
	fmt.Printf("Car %d: Arriving\n", id)
	p.spaces <- struct{}{}
	fmt.Printf("Car %d: Parking (Free spaces: %d)\n", id, p.capacity-len(p.spaces))
}

func (p *Park) Depart(id int) {
	fmt.Printf("Car %d: Departing\n", id)
	<-p.spaces
	fmt.Printf("Car %d: Left (Free spaces: %d)\n", id, p.capacity-len(p.spaces))
}

func (p *Park) Park(id int) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func Car(id int, p *Park, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		p.Arrive(id)
		p.Park(id)
		p.Depart(id)
	}
}

func main() {
	const (
		capacity = 4
		cars     = 20
	)

	park := NewPark(capacity)
	var wg sync.WaitGroup

	for i := 0; i < cars; i++ {
		wg.Add(1)
		go Car(i, park, &wg)
	}

	wg.Wait()
	fmt.Println("All cars have finished parking.")
}
