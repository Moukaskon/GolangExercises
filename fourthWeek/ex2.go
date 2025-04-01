// Η ώρα για επανάληψη έφτασε! Εντάξει τώρα νιώθω καλύτερα με τα channels. 
// Δεν βρήκα κάποια άλλη δομή, αν έχετε κάποια πρόταση θα την ακούσω με χαρά.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	content string
}

type Producer struct {
	queue chan Message
}

func (p *Producer) Start() {
	for i := 0; i < 10; i++ {
		msg := Message{content: fmt.Sprintf("Message %d", i)}
		fmt.Println("Produced:", msg.content)
		p.queue <- msg
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
	close(p.queue)
}

type Consumer struct {
	queue chan Message
}

func (c *Consumer) Start() {
	for msg := range c.queue {
		fmt.Println("Consumed:", msg.content)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
	fmt.Println("Consumer finished.")
}

func main() {
	queue := make(chan Message, 10)

	producer := &Producer{queue: queue}
	consumer := &Consumer{queue: queue}

	go producer.Start()
	go consumer.Start()

	time.Sleep(5 * time.Second)
}
