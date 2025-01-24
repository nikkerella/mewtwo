// Implement a set-like structure in Go.
// Since struct{} takes zero memory,
// it is often used as the value type in a map.
package main

import "fmt"

func main() {
	set := make(map[string]struct{})

	set["pikachu"] = struct{}{}
	set["raichu"] = struct{}{}

	// Check duplicate
	if _, exists := set["raichu"]; exists {
		fmt.Println("Raichu exists in the set!")
	}

	delete(set, "raichu")
}
