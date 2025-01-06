```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        var mu sync.Mutex // Add mutex for synchronization
        const numWorkers = 10

        for i := 0; i < numWorkers; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        mu.Lock() // Acquire lock before accessing count
                        count++
                        mu.Unlock() // Release lock
                }()
        }

        wg.Wait()
        fmt.Println("Count:", count)
}
```