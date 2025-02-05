// An empty struct can be used in channels
// when you only care about signaling and not the data being passed.
// It's useful for implementing notifications or simple synchronization.
package main

import (
	"fmt"
	"time"
)

func worker(done chan struct{}) {
	fmt.Println("Working...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done!")
	done <- struct{}{} // Signal that the work is done
}

func main() {
	done := make(chan struct{})

	go worker(done)

	// Wait for the worker to finish
	<-done
	fmt.Println("Worker has finished.")
}
