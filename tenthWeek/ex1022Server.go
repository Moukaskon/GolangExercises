// Τώρα εμφανίζετε το όνομα του πελάτη όταν συνδέεται και αποσυνδέεται από τον διακομιστή,
// καθώς και όταν στέλνει μηνύματα. Αρκετά ενδιαφέρουσα εργασία! Μου άρεσε αυτή η προσθήκη.
// Νιώθω ότι κάπως έτσι θα ήταν στα 80s τα πράγματα χαχαχα

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	clients  = make(map[net.Conn]string)
	mutex    = sync.Mutex{}
	clientID = 0
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
		clientID++
		name := fmt.Sprintf("Client#%d", clientID)
		clients[conn] = name
		mutex.Unlock()

		broadcast(fmt.Sprintf("%s has joined the chat.\n", name), conn)

		go handleClient(conn, name)
	}
}

func handleClient(conn net.Conn, name string) {
	defer func() {
		mutex.Lock()
		delete(clients, conn)
		mutex.Unlock()

		conn.Close()
		broadcast(fmt.Sprintf("%s has left the chat.\n", name), conn)
	}()

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.TrimSpace(msg)
		if msg == "CLOSE" {
			return
		}

		broadcast(fmt.Sprintf("%s: %s\n", name, msg), conn)
	}
}

func broadcast(message string, sender net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()

	for conn := range clients {
		if conn != sender {
			fmt.Fprint(conn, message)
		}
	}
}
