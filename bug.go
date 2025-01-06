```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        const numWorkers = 10

        for i := 0; i < numWorkers; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        count++
                }()
        }

        wg.Wait()
        fmt.Println("Count:", count)
}
```