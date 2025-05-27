package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:1234")
    if err != nil {
        fmt.Println("Unable to connect:", err)
        return
    }
    defer conn.Close()

    fmt.Println("Connected to server.")

    go receiveMessages(conn)
    sendMessages(conn)
}

func sendMessages(conn net.Conn) {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("> ")
        if !scanner.Scan() {
            break
        }
        text := scanner.Text()
        fmt.Fprintln(conn, text)
        if text == "CLOSE" {
            return
        }
    }
}

func receiveMessages(conn net.Conn) {
    reader := bufio.NewReader(conn)
    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            return
        }
        fmt.Print(msg)
    }
}
