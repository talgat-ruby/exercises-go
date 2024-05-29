package problem12

import (
	"fmt"
	"sort"
)

func keysAndValues[T comparable, S comparable](input map[T]S) ([]T, []S) {
	keys := make([]T, 0, len(input))
	values := make([]S, len(input))

	for key := range input {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprintf("%v", keys[i]) < fmt.Sprintf("%v", keys[j])
	})

	for i, key := range keys {
		values[i] = input[key]
	}

	return keys, values
}
