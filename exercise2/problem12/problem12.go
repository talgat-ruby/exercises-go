package main

import (
	"fmt"
	"sort"
)

// keysAndValues takes a map and returns two slices: sorted keys and their corresponding values.
func keysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	// Collect keys from the map
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// Sort keys
	sort.Slice(keys, func(i, j int) bool {
		switch k := any(keys[i]).(type) {
		case string:
			return k < keys[j].(string)
		case int:
			return k < keys[j].(int)
		default:
			return fmt.Sprintf("%v", keys[i]) < fmt.Sprintf("%v", keys[j])
		}
	})

	// Collect values in the
