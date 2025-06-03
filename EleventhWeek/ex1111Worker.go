package main

import (
    "encoding/binary"
    "net"
    "runtime"
    "sync"
)

func computePartialPI(n, id, totalWorkers int) float64 {
    h := 1.0 / float64(n)
    var indices []int
    for i := id; i < n; i += totalWorkers {
        indices = append(indices, i)
    }

    numThreads := runtime.NumCPU()
    chunkSize := len(indices) / numThreads
    if chunkSize == 0 {
        chunkSize = 1
    }

    results := make(chan float64, numThreads)
    var wg sync.WaitGroup

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

    var total float64
    for r := range results {
        total += r
    }

    return h * total
}

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    var n32, id32, total32 int32
    binary.Read(conn, binary.BigEndian, &n32)
    binary.Read(conn, binary.BigEndian, &id32)
    binary.Read(conn, binary.BigEndian, &total32)

    result := computePartialPI(int(n32), int(id32), int(total32))

    // Send result: float64
    binary.Write(conn, binary.BigEndian, result)
}
