package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter struct with a mutex lock
type SafeCounter struct {
	mu  sync.Mutex
	val int
}

// Increment method with mutex locking
func (c *SafeCounter) Increment() {
	c.mu.Lock() // Lock before modifying shared resource
	c.val++
	c.mu.Unlock() // Unlock after modification
}

// GetValue method to safely read the value
func (c *SafeCounter) GetValue() int {
	c.mu.Lock() // Lock before reading shared resource
	defer c.mu.Unlock()
	return c.val
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	// Start 5 goroutines that increment the counter
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				counter.Increment()
				time.Sleep(time.Millisecond * 10) // Simulate some work
			}
			wg.Done()
		}()
	}

	wg.Wait() // Wait for all goroutines to finish

	// Print the final counter value
	fmt.Println("Final Counter Value:", counter.GetValue())
}
