package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("Server error:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Calculator server started on port 12345")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go handleCalculator(conn)
	}
}

func handleCalculator(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	line = strings.TrimSpace(line)
	parts := strings.Split(line, "|")
	if len(parts) != 3 {
		conn.Write([]byte("ERROR: Invalid format\n"))
		return
	}

	op := parts[0]
	a, err1 := strconv.Atoi(parts[1])
	b, err2 := strconv.Atoi(parts[2])
	if err1 != nil || err2 != nil {
		conn.Write([]byte("ERROR: Invalid operands\n"))
		return
	}

	var result int
	var errMsg string

	switch op {
	case "ADD":
		result = a + b
	case "SUB":
		result = a - b
	case "MUL":
		result = a * b
	case "DIV":
		if b == 0 {
			errMsg = "ERROR: Division by zero"
		} else {
			result = a / b
		}
	default:
		errMsg = "ERROR: Unknown operation"
	}

	if errMsg != "" {
		conn.Write([]byte(errMsg + "\n"))
	} else {
		conn.Write([]byte(fmt.Sprintf("RESULT: %d\n", result)))
	}
}
