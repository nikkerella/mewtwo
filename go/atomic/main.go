package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	// Create an atomic.Value
	var val atomic.Value

	// Store a value in the atomic.Value
	val.Store("Skylar")
	// val.Store(39)

	// Load the value from the atomic.Value
	loadedValue := val.Load()

	// Type assert the strVal value to a string
	strVal := loadedValue.(string)
	// intVal := loadedValue.(int)

	fmt.Println("Loaded string value:", strVal)
	// fmt.Println("Loaded int value:", intVal)
}
