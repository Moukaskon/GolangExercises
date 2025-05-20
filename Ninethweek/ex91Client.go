// Τρέχετε πρώτα τον server και μετά τον client με την εντολή go run ex91Server.go και go run ex91Client.go αντίστοιχα.
// Ο client θα συνδεθεί στον server και θα στείλει το αίτημα για μετατροπή κειμένου. Ο client κλείνει μετά την ολοκλήρωση της διαδικασίας,
// ενώ ο server παραμένει σε λειτουργία και περιμένει για νέες συνδέσεις.

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

	fmt.Println("Choose operation:")
	fmt.Println("1. Convert UPPERCASE to lowercase")
	fmt.Println("2. Convert lowercase to UPPERCASE")
	fmt.Println("3. Encode with Caesar Cipher")
	fmt.Println("4. Decode with Caesar Cipher")
	fmt.Print("Enter choice (1-4): ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var op string
	var key string = "0"

	switch choice {
	case "1":
		op = "LOW"
	case "2":
		op = "UPP"
	case "3":
		op = "ENC"
		fmt.Print("Enter Caesar cipher key: ")
		key, _ = reader.ReadString('\n')
		key = strings.TrimSpace(key)
	case "4":
		op = "DEC"
		fmt.Print("Enter Caesar cipher key: ")
		key, _ = reader.ReadString('\n')
		key = strings.TrimSpace(key)
	default:
		fmt.Println("Invalid choice")
		return
	}

	fmt.Print("Enter message: ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	request := fmt.Sprintf("%s|%s|%s\n", op, key, message)
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println("Response from server:", strings.TrimSpace(response))
}
