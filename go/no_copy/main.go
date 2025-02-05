package main

import (
	"fmt"
	"sync"
)

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type Counter struct {
	noCopy noCopy
	mu     sync.Mutex
	count  int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	original := Counter{}

	// This creates a copy of the counter
	copy := original

	// Using the copy to increment
	copy.Increment()

	// Original counter's value remains unchanged
	fmt.Println("Original value:", original.Get()) // Output: 0
	fmt.Println("Copy value:", copy.Get())         // Output: 1
}
