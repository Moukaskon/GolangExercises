// Αυτός ήταν ο τρόπος που βρήκα χωρίς το gob, δεν είδα κάτι πιο αποτελεσματικό απο αυτό.
package main

import (
    "encoding/binary"
    "fmt"
    "net"
)

func main() {
    const N = 4
    const n = 100000

    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    defer ln.Close()

    fmt.Println("Master waiting for workers...")

    results := make(chan float64, N)

    for i := 0; i < N; i++ {
        conn, err := ln.Accept()
        if err != nil {
            panic(err)
        }

        go func(conn net.Conn, id int) {
            defer conn.Close()

            binary.Write(conn, binary.BigEndian, int32(n))
            binary.Write(conn, binary.BigEndian, int32(id))
            binary.Write(conn, binary.BigEndian, int32(N))

            var partial float64
            binary.Read(conn, binary.BigEndian, &partial)
            results <- partial
        }(conn, i)
    }

    var pi float64
    for i := 0; i < N; i++ {
        pi += <-results
    }

    fmt.Printf("Final approximation of Pi: %.12f\n", pi)
}
