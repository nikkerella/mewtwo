package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// You can create an errgroup.Group using errgroup.WithContext,
	// which also provides a context.Context that will be canceled
	// if any goroutine returns an error.
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Goroutine 1 done")
			return nil
		case <-ctx.Done():
			fmt.Println("Goroutine 1 canceled")
			return ctx.Err()
		}
	})

	g.Go(func() error {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Goroutine 2 done")
			// This will cancel the group
			return errors.New("error in goroutine 2")
		case <-ctx.Done():
			fmt.Println("Goroutine 2 canceled")
			return ctx.Err()
		}
	})

	g.Go(func() error {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Goroutine 3 done")
			return nil
		case <-ctx.Done():
			fmt.Println("Goroutine 3 canceled")
			return ctx.Err()
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("All goroutines completed successfully")
	}
}
