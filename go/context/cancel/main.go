package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a cancelCtx (using context.WithCancel)
	ctx, cancel := context.WithCancel(context.Background())

	// Start a goroutine that waits for the context to be canceled
	go func() {
		<-ctx.Done() // Wait for the `done` channel to be closed
		fmt.Println("Context canceled")
	}()

	// Simulate cancellation after 1 second
	time.Sleep(1 * time.Second)
	cancel()

	// Wait for the goroutine to finish
	time.Sleep(1 * time.Second)
}
