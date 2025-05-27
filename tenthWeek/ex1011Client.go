package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close()

    fmt.Print("Enter n (e.g. 100000) or -1 to exit: ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')

    conn.Write([]byte(input))

    reply := bufio.NewReader(conn)
    response, _ := reply.ReadString('\n')
    fmt.Println("Server response:", response)
}
