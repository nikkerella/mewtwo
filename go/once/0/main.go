package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var once sync.Once
var value int = 999

func plus() {
	value += 1
	fmt.Println("Function executed")
}

func safePlus() {
	mu.Lock()
	defer mu.Unlock()

	value += 1
	fmt.Println("Safe function executed")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()

			// plus()
			once.Do(plus)

			// safePlus()
			// once.Do(safePlus)

			fmt.Println(value)
		}()
	}

	wg.Wait()
}
