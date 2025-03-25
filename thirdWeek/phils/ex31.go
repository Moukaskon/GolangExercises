package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Philosopher struct {
	identity int
	next     int
	scale    int
	left     *Fork
	right    *Fork
}

func NewPhilosopher(id, n, s int, l, r *Fork) *Philosopher {
	return &Philosopher{
		identity: id,
		next:     n,
		scale:    s,
		left:     l,
		right:    r,
	}
}

func (p *Philosopher) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Thinking
		fmt.Printf("Philosopher %d is thinking\n", p.identity)
		p.Delay(p.scale)

		// Hungry
		fmt.Printf("Philosopher %d is trying to get fork %d\n", p.identity, p.identity)
		p.right.Get()

		// Got right fork
		fmt.Printf("Philosopher %d got fork %d and is trying to get fork %d\n", p.identity, p.identity, p.next)
		p.left.Get()

		// Eating
		fmt.Printf("Philosopher %d is eating\n", p.identity)
		p.Delay(p.scale)

		// Releasing forks
		fmt.Printf("Philosopher %d is releasing left fork %d\n", p.identity, p.next)
		p.left.Put()

		fmt.Printf("Philosopher %d is releasing fork %d\n", p.identity, p.identity)
		p.right.Put()
	}
}

func (p *Philosopher) Delay(scale int) {
	time.Sleep(time.Duration(rand.Intn(scale)) * time.Millisecond)
}
