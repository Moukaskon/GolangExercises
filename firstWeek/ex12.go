package main

import (
	"fmt"
)

// Αν αυξήσουμε το αριθμό των επαναλήψεων θα δούμε μίξη των εκτυπώσεων.

func printNum(done chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	done <- true
}

func printHello(done chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello from print!")
	}
	done <- true
}

func main() {
	done := make(chan bool)
	go printNum(done)
	go printHello(done)

	<-done
	<-done
}