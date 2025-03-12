package main

import (
	"fmt"
)

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
	for i := 0; i < 10; i++ {
		go printNum(done)
		go printHello(done)
	}

	for i := 0; i < 20; i++ { // Περιμένουμε να τερματίσουν και τα 20 νήματα
		<-done
	}
}