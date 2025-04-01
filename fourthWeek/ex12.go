// Εδώ είναι η λύση για να μην έχουμε αδιέξοδο. Επίσης σας ευχαριστώ για την λύση με τους φιλόσοφους,
// είχε πράγματι ενδιαφέρον. Η αλήθεια είναι ότο μου πήρε λίγη ώρα να την καταλάβω,
// μάλλον χρειάζομαι επανάληψη σε channels...

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Buffer struct
type Buffer struct {
	contents int
	size     int
	counter  int
	lock     sync.Mutex
	full     *sync.Cond
	empty    *sync.Cond
	closed   bool
}

func NewBuffer(size int) *Buffer {
	b := &Buffer{size: size}
	b.full = sync.NewCond(&b.lock)
	b.empty = sync.NewCond(&b.lock)
	return b
}

func (b *Buffer) Put(data int) {
	b.lock.Lock()
	defer b.lock.Unlock()

	for b.counter == b.size {
		fmt.Println("The buffer is full")
		b.full.Wait()
	}

	b.contents = data
	b.counter++
	fmt.Printf("Prod %v No %d Count = %d\n", time.Now().UnixNano(), data, b.counter)

	b.empty.Signal()
}

func (b *Buffer) Get() (int, bool) {
	b.lock.Lock()
	defer b.lock.Unlock()

	for b.counter == 0 {
		if b.closed {
			return 0, false
		}
		fmt.Println("The buffer is empty")
		b.empty.Wait()
	}

	data := b.contents
	b.counter--
	fmt.Printf("  Cons %v No %d Count = %d\n", time.Now().UnixNano(), data, b.counter)

	b.full.Signal()
	return data, true
}

func (b *Buffer) CloseBuffer() {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.closed = true
	b.empty.Broadcast()
}

type Consumer struct {
	buff  *Buffer
	scale int
}

func (c *Consumer) Start(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		time.Sleep(time.Duration(rand.Intn(c.scale)) * time.Millisecond)
		if _, ok := c.buff.Get(); !ok {
			fmt.Println("Consumer exiting...")
			return
		}
	}
}

type Producer struct {
	buff  *Buffer
	reps  int
	scale int
}

func (p *Producer) Start(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < p.reps; i++ {
		p.buff.Put(i)
		time.Sleep(time.Duration(rand.Intn(p.scale)) * time.Millisecond)
	}
	p.buff.CloseBuffer()
	fmt.Println("Producer finished")
}

func main() {
	bufferSize := 1
	noIterations := 20
	producerDelay := 100
	consumerDelay := 1

	buff := NewBuffer(bufferSize)

	var wg sync.WaitGroup

	wg.Add(2)
	go (&Producer{buff, noIterations, producerDelay}).Start(&wg)
	go (&Consumer{buff, consumerDelay}).Start(&wg)

	wg.Wait()
	fmt.Println("All done!")
}
