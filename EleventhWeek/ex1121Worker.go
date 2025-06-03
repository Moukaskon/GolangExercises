package main

import (
    "encoding/gob"
    "net"
)

type Task struct {
    N  int
    ID int
    TotalWorkers int
}

func computePI(n, id, totalWorkers int) float64 {
    var sum float64
    h := 1.0 / float64(n)
    for i := id; i < n; i += totalWorkers {
        x := h * (float64(i) + 0.5)
        sum += 4.0 / (1.0 + x*x)
    }
    return h * sum
}

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    dec := gob.NewDecoder(conn)
    enc := gob.NewEncoder(conn)

    var task Task
    err = dec.Decode(&task)
    if err != nil {
        panic(err)
    }

    result := computePI(task.N, task.ID, task.TotalWorkers)
    enc.Encode(result)
}
