# Go Race Condition Example

This repository demonstrates a common race condition in Go that occurs when multiple goroutines concurrently access and modify a shared variable without proper synchronization.

The `bug.go` file shows the problematic code where a counter is incremented by multiple goroutines without a mutex to protect it.  The `bugSolution.go` file provides a corrected version using a mutex to ensure thread safety. 

## How to run

1. Clone the repository
2. Navigate to the project directory
3. Run the problematic code: `go run bug.go` 
4. Run the corrected code: `go run bugSolution.go`

Observe the different outputs to understand the impact of the race condition and the solution.