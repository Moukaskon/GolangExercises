package main

import "sync"

type Fork struct {
	identity int
	lock     sync.Mutex
}

func NewFork(id int) *Fork {
	return &Fork{identity: id}
}

func (f *Fork) Get() {
	f.lock.Lock()
}

func (f *Fork) Put() {
	f.lock.Unlock()
}
