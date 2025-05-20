package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		fmt.Println("Could not connect to server:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Calculator Options:")
	fmt.Println("1. ADD (Addition)")
	fmt.Println("2. SUB (Subtraction)")
	fmt.Println("3. MUL (Multiplication)")
	fmt.Println("4. DIV (Division)")
	fmt.Print("Choose operation (1-4): ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var op string
	switch choice {
	case "1":
		op = "ADD"
	case "2":
		op = "SUB"
	case "3":
		op = "MUL"
	case "4":
		op = "DIV"
	default:
		fmt.Println("Invalid option")
		return
	}

	fmt.Print("Enter first operand: ")
	a, _ := reader.ReadString('\n')
	a = strings.TrimSpace(a)

	fmt.Print("Enter second operand: ")
	b, _ := reader.ReadString('\n')
	b = strings.TrimSpace(b)

	request := fmt.Sprintf("%s|%s|%s\n", op, a, b)
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Println("Send error:", err)
		return
	}

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Receive error:", err)
		return
	}
	fmt.Println("Server says:", strings.TrimSpace(response))
}
