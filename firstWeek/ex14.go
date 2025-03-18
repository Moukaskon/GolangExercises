package main

import (
	"fmt"
)

func printPol(k int, done chan bool) {
	for i := 1; i < 21; i++ {
		fmt.Println(i * k)
	}
	done <- true
}

func main() {
	done := make(chan bool) // Θα μπορούσαμε να χρησιμοποιήσουμε buffered channel για μεγαλύτερη ταχύτητα.
	for i := 0; i < 10; i++ {
		go printPol(i+1, done)
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}