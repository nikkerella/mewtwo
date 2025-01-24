package main

import (
	"fmt"
	"sync"
)

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type Counter struct {
	noCopy noCopy // 標記結構體不可複製
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
	honmono := Counter{}

	// This creates a nisemono of the counter
	nisemono := honmono

	// Using the copy to increment
	nisemono.Increment()

	// Original counter's value remains unchanged
	fmt.Println("Original value:", honmono.Get()) // Output: 0
	fmt.Println("Copy value:", nisemono.Get())    // Output: 1
}
