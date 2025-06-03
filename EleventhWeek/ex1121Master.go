// Σας στέλνω και τα built αρχεία για να τα τρέξετε, αφού ανοίξετε έναν master ανοίγετε Ν (4) workers
// και μετα θα εμφανιστεί το αποτέλεσμα του υπολογισμού του π. 

package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Task struct {
	N            int
	ID           int
	TotalWorkers int
}

func main() {
	const N = 4 // number of workers
	const n = 100000

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	results := make(chan float64, N)

	for i := 0; i < N; i++ {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn, id int) {
			defer conn.Close()

			enc := gob.NewEncoder(conn)
			dec := gob.NewDecoder(conn)

			task := Task{N: n, ID: id, TotalWorkers: N}
			enc.Encode(task)

			var partial float64
			dec.Decode(&partial)

			results <- partial
		}(conn, i)
	}

	var pi float64
	for i := 0; i < N; i++ {
		pi += <-results
	}

	fmt.Printf("Approximation of Pi: %.12f\n", pi)
}
