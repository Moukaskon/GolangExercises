package main

import (
    "encoding/gob"
    "net"
    "runtime"
    "sync"
)

type Task struct {
    N            int
    ID           int
    TotalWorkers int
}

func computePartialPI(n, id, totalWorkers int) float64 {
    numThreads := runtime.NumCPU()
    var wg sync.WaitGroup
    results := make(chan float64, numThreads)

    h := 1.0 / float64(n)
    var indices []int
    for i := id; i < n; i += totalWorkers {
        indices = append(indices, i)
    }

    chunkSize := len(indices) / numThreads
    if chunkSize == 0 {
        chunkSize = 1
    }

    for t := 0; t < numThreads; t++ {
        start := t * chunkSize
        end := start + chunkSize
        if t == numThreads-1 {
            end = len(indices)
        }

        wg.Add(1)
        go func(start, end int) {
            defer wg.Done()
            var localSum float64
            for i := start; i < end; i++ {
                x := h * (float64(indices[i]) + 0.5)
                localSum += 4.0 / (1.0 + x*x)
            }
            results <- localSum
        }(start, end)
    }

    wg.Wait()
    close(results)

    var totalSum float64
    for r := range results {
        totalSum += r
    }

    return h * totalSum
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
    if err := dec.Decode(&task); err != nil {
        panic(err)
    }

    result := computePartialPI(task.N, task.ID, task.TotalWorkers)
    enc.Encode(result)
}
