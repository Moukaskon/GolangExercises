// Τα αποτελέσματα αποθηκεύονται οπότε έχουμε πολύ γρήγορη απάντηση τη 2η φορά που ζητηθεί το ίδιο αποτέλεσμα.
// Για δοκιμή, το 1000000000 την πρώτη φορά κάνει περίπου 2 δευτερόλεπτα, αλλά τη δεύτερη φορά απαντάει αμέσως.

package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
)

var (
	piCache = make(map[int]float64)
	mutex   = sync.Mutex{}
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Server is running on port 8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read:", err)
		return
	}

	message = strings.TrimSpace(message)
	n, err := strconv.Atoi(message)
	if err != nil {
		conn.Write([]byte("Invalid number\n"))
		return
	}

	if n == -1 {
		conn.Write([]byte("Goodbye\n"))
		return
	}

	var pi float64

	// Check cache with mutex protection
	mutex.Lock()
	val, exists := piCache[n]
	mutex.Unlock()

	if exists {
		pi = val
	} else {
		pi = calculatePi(n)
		mutex.Lock()
		piCache[n] = pi
		mutex.Unlock()
	}

	conn.Write([]byte(fmt.Sprintf("PI approximation: %.10f\n", pi)))
}

func calculatePi(n int) float64 {
	sum := 0.0
	h := 1.0 / float64(n)
	for i := 0; i < n; i++ {
		x := h * (float64(i) + 0.5)
		sum += 4.0 / (1.0 + x*x)
	}
	return h * sum
}
