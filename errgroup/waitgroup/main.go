package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	errCh := make(chan error, 1) // A channel to catch errors
	var err error                // The error caught

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Goroutine 1 done")
		case <-ctx.Done():
			fmt.Println("Goroutine 1 canceled:", ctx.Err())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-time.After(1 * time.Second):
			// Simulate an error
			errCh <- fmt.Errorf("error in goroutine 2")
		case <-ctx.Done():
			fmt.Println("Goroutine 2 canceled:", ctx.Err())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-time.After(3 * time.Second):
			// Simulate an error
			errCh <- fmt.Errorf("error in goroutine 3")
		case <-ctx.Done():
			fmt.Println("Goroutine 3 canceled:", ctx.Err())
		}
	}()

	// Monitor errors from goroutines
	go func() {
		err = <-errCh
		if err != nil {
			fmt.Println("Received error:", err)
			cancel() // Cancel the context to stop other goroutines
		}
	}()

	wg.Wait()

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("All goroutines completed successfully")
	}
}
