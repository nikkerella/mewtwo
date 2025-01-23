package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

// Cache struct
type Cache struct {
	data    string
	expired bool
	mu      sync.RWMutex
	sfGroup singleflight.Group
}

// To simulate cache regeneration
func (c *Cache) regenerateCache() string {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Simulate long cache regeneration
	time.Sleep(3 * time.Second)

	c.data = "데이터"
	c.expired = false

	fmt.Println("Regenerated cache!")

	return c.data
}

// To simulate a user requesting the data
func (c *Cache) requestData(userID int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquire read lock
	c.mu.RLock()

	// Simulate the cache being expired
	if c.expired {
		c.mu.RUnlock()

		fmt.Println("UserID", userID, "Cache expired. Requesting regeneration...")

		c.regenerateCache()

		fmt.Println("UserID", userID, "Cache regenerated. The data:", c.data)
	} else {
		fmt.Println("UserID", userID, "Cache is available. Fetching data:", c.data)
		c.mu.RUnlock()
	}
}

func (c *Cache) singleflight(userID int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquire read lock
	c.mu.RLock()

	// Simulate the cache being expired
	if c.expired {
		c.mu.RUnlock()

		fmt.Println("UserID", userID, "Cache expired. Requesting regeneration...")
		// Use singleflight to ensure only one goroutine regenerates the cache
		result, _, _ := c.sfGroup.Do("regenerateCache", func() (interface{}, error) {
			return c.regenerateCache(), nil
		})

		// Retrieve regenerated data
		c.data = result.(string)

		fmt.Println("UserID", userID, "Cache regenerated. The data:", c.data)
	} else {
		fmt.Println("UserID", userID, "Cache is available. Fetching data:", c.data)
		c.mu.RUnlock()
	}
}

func main() {
	cache := &Cache{
		data:    "データ",
		expired: true,
	}

	users := 3
	var wg sync.WaitGroup

	// Simulate cache stampede
	if true {
		for i := 1; i <= users; i++ {
			wg.Add(1)
			go cache.requestData(i, &wg)
		}
	}

	// Test singleflight
	if false {
		for i := 1; i <= users; i++ {
			wg.Add(1)
			go cache.singleflight(i, &wg)
		}
	}

	wg.Wait()
}
