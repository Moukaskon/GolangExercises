package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started on port 12345")

	for {
		conn, err := listener.Accept()
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
	request, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	request = strings.TrimSpace(request)
	parts := strings.SplitN(request, "|", 3)
	if len(parts) != 3 {
		conn.Write([]byte("ERROR: Invalid request format\n"))
		return
	}

	op := parts[0]
	key, _ := strconv.Atoi(parts[1])
	message := parts[2]
	var response string

	switch op {
	case "LOW":
		response = toLower(message)
	case "UPP":
		response = toUpper(message)
	case "ENC":
		response = caesarCipher(message, key)
	case "DEC":
		response = caesarCipher(message, -key)
	default:
		response = "ERROR: Unknown operation"
	}

	conn.Write([]byte(response + "\n"))
}

func toLower(msg string) string {
	var sb strings.Builder
	for _, r := range msg {
		if unicode.IsUpper(r) {
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func toUpper(msg string) string {
	var sb strings.Builder
	for _, r := range msg {
		if unicode.IsLower(r) {
			sb.WriteRune(unicode.ToUpper(r))
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func caesarCipher(msg string, offset int) string {
	var sb strings.Builder
	for _, r := range msg {
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			newChar := 'a' + (r-'a'+rune((offset%26)+26))%26
			sb.WriteRune(newChar)
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
