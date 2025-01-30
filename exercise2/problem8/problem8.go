package main

import "fmt"

// simplify maps each element of the input slice to its index and returns the map
func simplify(elements []string) map[string]int {
	result := make(map[string]int)

	// load function that stores the index of each element in the map
	load := func(key string, index int) {
		result[key] = index
	}

	// Iterate over elements and load them into the map
	for i, elem := range elements {
		load(elem, i)
	}

	return result
}

func main() {
	fmt.Println(simplify([]string{"a", "b", "c"}))               // Output: map[a:0 b:1 c:2]
	fmt.Println(simplify([]string{"z", "y", "x", "u", "v"}))     // Output: map[u:3 v:4 x:2 y:1 z:0]
}
