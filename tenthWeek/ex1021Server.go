// Ένας απλός TCP server που δέχεται συνδέσεις από πολλούς πελάτες και στέλνει μηνύματα σε όλους εκτός από τον αποστολέα.

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	clients = make(map[net.Conn]bool)
	mutex   = sync.Mutex{}
)

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
	fmt.Println("Server is running on port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}

		mutex.Lock()
		clients[conn] = true
		mutex.Unlock()

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer func() {
		mutex.Lock()
		delete(clients, conn)
		mutex.Unlock()
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return // Client disconnected
		}

		msg = strings.TrimSpace(msg)
		if msg == "CLOSE" {
			return
		}

		fmt.Printf("Received: %s\n", msg)
		broadcast(msg, conn)
	}
}

func broadcast(message string, sender net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()

	for conn := range clients {
		if conn != sender {
			fmt.Fprintf(conn, "Message from another client: %s\n", message)
		}
	}
}
